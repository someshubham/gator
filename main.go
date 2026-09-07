package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/someshubham/gator/internal/config"
	"github.com/someshubham/gator/internal/database"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Unable to read config")
		return
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println("Unable to open DB")
		return
	}

	dbQueries := database.New(db)

	s := state{
		config: &cfg,
		db:     dbQueries,
	}
	cmdList := commands{
		cmd: make(map[string]func(*state, command) error),
	}
	cmdList.register("login", handlerLogin)
	cmdList.register("register", handlerRegister)
	cmdList.register("reset", handlerReset)

	err = cmdList.run(&s, purifyArgs(os.Args))
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
