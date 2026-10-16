package tcphandler

import (
	"fmt"
	"os"
)

func Serve() string {
	return "HTTP/1.0 200 OK\nContent-Type: text/html\nContent-Length: 47\n\n<html><body>Hello from my server!</body></html>"
}

func ServeImage() string {
	file, err := os.ReadFile("/Users/sujas/Desktop/Tech/Personal/http-go/public/cat.jpg")
	if err != nil {
		fmt.Println("Error Reading File")
	}
	content_length := len(file)
	return fmt.Sprintf("HTTP/1.0 200 OK\nContent-Type: image/jpeg\nContent-Length: %d\n\n%s", content_length, file)
}

func ServeBinary() string {
	file, err := os.ReadFile("/Users/sujas/Desktop/Tech/Personal/http-go/public/http-go")
	if err != nil {
		fmt.Println("Error Reading File")
	}
	content_length := len(file)
	return fmt.Sprintf("HTTP/1.0 200 OK\nContent-Type: application/octet-stream\nContent-Length: %d\n\n%s", content_length, file)
}

func ServeError() string {
	return "HTTP/1.0 500 Internal Server Error\nContent-Type: application/json\nContent-Length: 35\n\n{'message':'Internal Server Error'}"
}
