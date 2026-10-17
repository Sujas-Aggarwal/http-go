package api

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Sujas-Aggarwal/http-go/internal/helper"
	"github.com/Sujas-Aggarwal/http-go/internal/tcp"
)

// Handler runs when a request matches a method and route.
type Handler func(*Request, *Response)

type Request struct {
	Headers helper.RequestHeader
	Body    []byte
}

type Response struct {
	conn   net.Conn
	method string
}

// Send writes one plain-text HTTP/1.0 response.
func (r *Response) Send(status int, body string) error {
	data := []byte(body)
	header := "HTTP/1.0 " + strconv.Itoa(status) + " " + statusText(status) + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"Content-Length: " + strconv.Itoa(len(data)) + "\r\n\r\n"
	if _, err := r.conn.Write([]byte(header)); err != nil {
		return err
	}
	if r.method == "HEAD" {
		return nil
	}
	_, err := r.conn.Write(data)
	return err
}

func statusText(status int) string {
	switch status {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Status"
	}
}

type App struct {
	port   string
	routes map[string]map[string]Handler
}

func InitApp(port string) *App {
	return &App{port: port, routes: make(map[string]map[string]Handler)}
}

func (a *App) Handle(method, route string, handler Handler) {
	if a.routes[method] == nil {
		a.routes[method] = make(map[string]Handler)
	}
	a.routes[method][route] = handler
}

func (a *App) Get(route string, handler Handler)  { a.Handle("GET", route, handler) }
func (a *App) Post(route string, handler Handler) { a.Handle("POST", route, handler) }
func (a *App) Head(route string, handler Handler) { a.Handle("HEAD", route, handler) }

// Listen accepts connections until the listener fails.
func (a *App) Listen() error {
	listener, err := net.Listen("tcp", a.port)
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Println("INFO: Successfully Started Listening on", a.port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go tcp.HandleConnectionWithHandler(conn, a.dispatch)
	}
}

func (a *App) dispatch(header helper.RequestHeader, body []byte, conn net.Conn) {
	response := &Response{conn: conn, method: header.Method}
	if handler := a.routes[header.Method][header.Route]; handler != nil {
		handler(&Request{Headers: header, Body: body}, response)
		return
	}
	_ = response.Send(404, "Not Found")
}
