package middlewares

import (
	"net/http"
)

type Middlewarer interface {
	Call(env Env) MiddlewareResponder
}

type Middleware struct {
	Next Middleware
}

type Env struct {
	W http.ResponseWriter
	R *http.Request
}

type MiddlewareResponder struct {
	Code        int
	ContentType string
	Content     string
}
