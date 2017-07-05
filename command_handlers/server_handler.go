package command_handlers

import (
	"../server"
)

func ServerStart(host, port string) {
	new_server := server.NewServer(host, port)
	new_server.Run()
}
