package commands

import (
	"../command_handlers"
	"../config"
	"../utils"
	"flag"
	"fmt"
	"os"
)

var AVAILABLE_COMMANDS = []string{"new", "s"}

var db_flag string
var host_flag string
var port_flag string

func init() {
	flag.StringVar(&db_flag, "d", config.DEFAULT_DATABASE, "specific default database")
	flag.StringVar(&host_flag, "h", config.DEFAULT_HOST, "specific default database")
	flag.StringVar(&port_flag, "p", config.DEFAULT_PORT, "specific default database")
	flag.Parse()
}

func CheckCommand() {
	if len(os.Args) < 2 || !utils.IsContainString(AVAILABLE_COMMANDS, os.Args[1]) {
		OutputHelpInfo()
		os.Exit(1)
	}
}

func OutputHelpInfo() {
	fmt.Println("usage: gota [command] [flag] [flag_value]")
	fmt.Println("available commands: ", AVAILABLE_COMMANDS)
}

func HandleCommand() {
	switch os.Args[1] {
	case "new":
		// command_handlers.NewProject()
	case "s":
		command_handlers.ServerStart(host_flag, port_flag)
	default:
		OutputHelpInfo()
		os.Exit(1)
	}

}
