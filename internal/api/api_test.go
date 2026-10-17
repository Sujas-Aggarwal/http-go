package api

import (
	"io"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/Sujas-Aggarwal/http-go/internal/tcp"
)

func TestRequestToResponse(t *testing.T) {
	app := InitApp(":0")
	app.Get("/hello", func(req *Request, res *Response) {
		_ = res.Send(200, req.Headers.Method+" "+req.Headers.Route)
	})
	app.Post("/echo", func(req *Request, res *Response) {
		_ = res.Send(201, string(req.Body))
	})

	tests := []struct {
		name   string
		chunks []string
		status string
		body   string
	}{
		{"GET", []string{"GET /hello HTTP/1.0\r\nHost: localhost\r\n\r\n"}, "HTTP/1.0 200 OK", "GET /hello"},
		{"POST split body", []string{"POST /echo HTTP/1.0\r\nContent-Length: 5\r\n\r\nhe", "llo"}, "HTTP/1.0 201 Created", "hello"},
		{"missing route", []string{"GET /missing HTTP/1.0\r\n\r\n"}, "HTTP/1.0 404 Not Found", "Not Found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			defer client.Close()
			if err := client.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
				t.Fatal(err)
			}
			go tcp.HandleConnectionWithHandler(server, app.dispatch)

			for _, chunk := range tt.chunks {
				if _, err := client.Write([]byte(chunk)); err != nil {
					t.Fatal(err)
				}
			}
			response, err := io.ReadAll(client)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.HasPrefix(string(response), tt.status) || !strings.HasSuffix(string(response), "\r\n\r\n"+tt.body) {
				t.Fatalf("unexpected response: %q", response)
			}
		})
	}
}
