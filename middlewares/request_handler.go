package middlewares

import (
	"fmt"
	"net/http"
)

type RequestHandler struct {
}

func (req *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
