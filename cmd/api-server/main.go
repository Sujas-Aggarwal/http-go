package main

import (
	"log"

	httpgo "github.com/Sujas-Aggarwal/http-go"
)

// this is gonna be a simple app which is actually gonna utilise the http-go library that we created and actually make something useful, makes sense?
func main() {
	app := httpgo.New(":8000")
	app.Get("/", func(req *httpgo.Request, res *httpgo.Response) {
		res.Send(200, "Hello from http-go! You requested "+req.Headers.Route)
	})
	app.Post("/echo", func(req *httpgo.Request, res *httpgo.Response) {
		res.Send(200, string(req.Body))
	})
	if err := app.Listen(); err != nil {
		log.Fatal(err)
	}
}
