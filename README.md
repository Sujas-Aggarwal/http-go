ok, for the time being, lets just implement http 1.0 for now,
i wanted to start from 0.9 but unfortunately, most browsers dont
support it now and we might not be able to test it properly.

so we will go with http 1.0 for now

specs -
pretty simple actually, since http 1.0 introduced support for media types also,
1.
first we need to specify the protocol we are using to send the data
2.
we need to send the content header to tell the type of the file we are sending
like text, html, css or images or something.
3.
then we also need to tell the size of the stuff we are sending in bytes

after these things, we can just directly send the stuff that we wanna send in what
ever format we specified in the content header type, maybe 8 bit encoded strings (utf-8)
or ascii or some weird ass format, doesn't matter tbh, depends on the end consumer.

since http1.0 works on tcp, we will also listen on tcp (i am not implementing tcp here)

# Issue 1 -
how do i know when to stop waiting for data and instead respond back?
## Solutoion - 
well easy, for GET requests sent by browser, the header always ends with \r\n\r\n
we need to use that


Now plan,
let's try to send some files over tcp/http instead of just plain text
- send binary files, http files, images and stuff.


i am thinking of exposing a very similar fashion of api as nodejs or somethign
- create a app, define a port you want to listen
- on that app, register routes and register middleware if you want
- then add a function that you want to execute on recieving that route and method
- in that funciton, you will have access to 2 vars, request and response
- the request one will provide you the info of the request, particularly the headers
- and the body will provide you with the body of the request

## Try the library

Import `github.com/Sujas-Aggarwal/http-go` from another Go module:

```go
package main

import (
	"log"

	httpgo "github.com/Sujas-Aggarwal/http-go"
)

func main() {
	app := httpgo.New(":8000")
	app.Get("/", func(req *httpgo.Request, res *httpgo.Response) {
		log.Println(req.Headers.Method, req.Headers.Route, req.Headers.Host)
		_ = res.Send(200, "Hello from http-go!")
	})
	app.Post("/echo", func(req *httpgo.Request, res *httpgo.Response) {
		_ = res.Send(200, string(req.Body))
	})
	log.Fatal(app.Listen())
}
```

You can also run `go run ./cmd/api-server` in this repository. Try
`curl localhost:8000/` or `curl -d hello localhost:8000/echo`.
