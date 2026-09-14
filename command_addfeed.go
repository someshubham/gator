package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/someshubham/gator/internal/database"
)

func handlerAddFeed(s *state, c command, usr database.User) error {
	if len(c.args) != 2 {
		fmt.Println("Required name and url of the feed")
		os.Exit(1)
	}
	feedName := c.args[0]
	url := c.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       url,
		UserID:    usr.ID,
	})

	if err != nil {
		fmt.Printf("Unable to create feed\n%s\n", err.Error())
		os.Exit(1)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		fmt.Printf("Unable to create feed follow\n%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(feed.Name)
	fmt.Println(feed.Url)

	return nil
}
