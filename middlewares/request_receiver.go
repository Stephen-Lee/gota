package middlewares

import (
	"fmt"
	"net/http"
)

type RequestReceiver struct {
}

func (req *RequestReceiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")

}
