package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/someshubham/gator/internal/database"
)

func handlerFollow(s *state, c command) error {

	if len(c.args) == 0 {
		fmt.Println("Need a url to run follow command")
		os.Exit(1)
	}

	usr, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		fmt.Printf("Unable to find the user\n%s\n", err.Error())
		os.Exit(1)
	}

	url := c.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		fmt.Printf("Unable to find the feed\n%s\n", err.Error())
		os.Exit(1)
	}

	followFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		fmt.Printf("Unable to find create feed follow\n%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(followFeed.FeedName)
	fmt.Println(followFeed.UserName)

	return nil
}
