package server

import (
	"../middlewares"
	"fmt"
	"log"
	"net/http"
)

type BaseServer struct {
	Host string
	Port string
}

func NewServer(host, port string) BaseServer {
	base_server := BaseServer{Host: host, Port: port}
	return base_server
}

func (bs *BaseServer) Run() {
	init_request_handler := &middlewares.RequestHandler{}
	bs.PrintServerInfo()
	listen_address := bs.Host + ":" + bs.Port
	log.Fatal(http.ListenAndServe(listen_address, init_request_handler))
}

func (bs *BaseServer) PrintServerInfo() {
	fmt.Println("host: ", bs.Host)
	fmt.Println("port: ", bs.Port)
}
