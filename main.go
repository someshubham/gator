package main

import (
	"fmt"
	"os"

	"github.com/someshubham/gator/internal/config"
)

func main() {

	s := state{
		config: &config.Config{
			DbUrl: "postgres://example",
		},
	}

	s.config.SetUser("")

	cmdList := commands{
		cmd: make(map[string]func(*state, command) error),
	}
	cmdList.register("login", handlerLogin)

	err := cmdList.run(&s, purifyArgs(os.Args))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func purifyArgs(args []string) command {
	if len(args) < 2 {
		fmt.Println("Expected args")
		os.Exit(1)
	}

	commandName := args[1]
	var commandArgs []string

	if len(args[2:]) == 0 {
		commandArgs = make([]string, 0)
	} else {
		commandArgs = args[2:]
	}

	return command{
		name: commandName,
		args: commandArgs,
	}
}
