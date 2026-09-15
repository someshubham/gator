package main

import (
	"context"
	"fmt"
	"os"

	"github.com/someshubham/gator/internal/database"
)

func handlerUnfollow(s *state, c command, usr database.User) error {

	if len(c.args) == 0 {
		fmt.Println("Need an input url for the unfollow")
		os.Exit(1)
	}

	url := c.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		fmt.Println("Unable to find the feed")
		os.Exit(1)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: usr.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		fmt.Printf("Unable to delete the feed %s \n", url)
		os.Exit(1)
	}

	return nil
}
