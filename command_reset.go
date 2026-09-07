package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, _ command) error {

	err := s.db.DeleteAllUsers(context.Background())

	if err != nil {
		fmt.Println("Unable to reset the user data")
		os.Exit(1)
	}

	return nil
}
