package main

import (
	"flag"
	"fmt"
	"github.com/utils"
	"os"
)

var AVAILABLE_COMMANDS = []string{"new", "help"}

var db_flag string
var help_flag string

func init() {
	flag.StringVar(&db_flag, "d", "mysql", "specific default database")
	flag.StringVar(&help_flag, "h", "", "help info")
	flag.Parse()
}

func CheckCommand() {
	if len(os.Args) < 2 || utils.IsContain(AVAILABLE_COMMANDS, os.Args[1]) {
		OutputHelpInfo()
		os.Exit(1)
	}
}

func OutputHelpInfo() {
	fmt.Println("usage: gota [command] [flag] [flag_value]")
	fmt.Println("available commands: ", AVAILABLE_COMMANDS)
}

func main() {
	CheckCommand()
	fmt.Println("server starting")
}
