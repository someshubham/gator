package main

import (
	"context"
	"fmt"
	"os"

	"github.com/someshubham/gator/internal/database"
)

func handlerFollowing(s *state, _ command, usr database.User) error {

	feedFollow, err := s.db.GetFeedFollowsForUser(context.Background(), usr.Name)

	if err != nil {
		fmt.Printf("Unable to fetch feed follows by user\n%s\n", err.Error())
		os.Exit(1)
	}

	for _, f := range feedFollow {
		fmt.Println(f.Name)
	}
	return nil
}
