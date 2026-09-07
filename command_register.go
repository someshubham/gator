package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/someshubham/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the name")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})

	if err != nil {
		fmt.Printf("Unable to add the user:\n%s", err.Error())
		os.Exit(1)
	}

	s.config.SetUser(user.Name)

	fmt.Println("User was created")
	fmt.Println(user)
	return nil
}
