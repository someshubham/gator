package main

import (
	"context"
	"fmt"
	"os"
)

func handlerFollowing(s *state, _ command) error {

	feedFollow, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)

	if err != nil {
		fmt.Printf("Unable to fetch feed follows by user\n%s\n", err.Error())
		os.Exit(1)
	}

	for _, f := range feedFollow {
		fmt.Println(f.Name)
	}
	return nil
}
