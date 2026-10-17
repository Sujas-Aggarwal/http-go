// Package httpgo provides a small HTTP/1.0 server API.
package httpgo

import "github.com/Sujas-Aggarwal/http-go/internal/api"

type App = api.App
type Handler = api.Handler
type Request = api.Request
type Response = api.Response

// New creates an app that will listen on an address such as ":8000".
func New(port string) *App { return api.InitApp(port) }
