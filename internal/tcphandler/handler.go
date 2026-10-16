package tcphandler

func Serve() string {
	return "HTTP/1.0 200 OK\nContent-Type: text/html\nContent-Length: 43\n\n<html><body>Hello from my server!</body></html>"
}
