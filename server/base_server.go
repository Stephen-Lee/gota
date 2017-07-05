package server

import (
	"github.com/valyala/fasthttp"
)

type BaseServer struct {
}

func Init() {
	addr = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
}

func Run() {
	// app := App.new()
	// run middleware
}
