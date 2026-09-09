package main

import (
	"context"
	"fmt"
	"os"
)

func handlerFeeds(s *state, _ command) error {

	feedsWithUser, err := s.db.GetFeedsWithUser(context.Background())

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	for _, feed := range feedsWithUser {
		fmt.Println(feed.FeedName)
		fmt.Println(feed.FeedUrl)
		fmt.Println(feed.UserName)
	}

	return nil
}
