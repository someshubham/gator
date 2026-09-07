package main

import (
	"context"
	"fmt"
	"os"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}

func (c *commands) run(s *state, command command) error {
	fn, ok := c.cmd[command.name]
	if ok {
		err := fn(s, command)
		if err != nil {
			return fmt.Errorf("Unable to run the command\n Error: %w", err)
		}
	} else {
		return fmt.Errorf("Unable to find the command")
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd[name] = f
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	usr, err := s.db.GetUser(context.Background(), cmd.args[0])

	if err != nil {
		fmt.Printf("Unable to login to unknown user\n%s", err.Error())
		os.Exit(1)
	}

	err = s.config.SetUser(usr.Name)

	if err != nil {
		return err
	}

	fmt.Printf("user %s has been set\n", s.config.CurrentUserName)
	return nil
}
