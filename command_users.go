package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func handlerUsers(s *state, _ command) error {
	usrs, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Unable to fetch any user\n%s\n", err.Error())
		os.Exit(1)
	}

	if len(usrs) == 0 {
		fmt.Println("No User available")
		os.Exit(1)
	}

	for _, usr := range usrs {
		isCurrentUser := strings.Compare(usr.Name, s.config.CurrentUserName) == 0

		if !isCurrentUser {
			fmt.Printf("* %s\n", usr.Name)
		} else {
			fmt.Printf("* %s (current)\n", usr.Name)
		}
	}

	return nil
}
