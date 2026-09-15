package main

import (
	"context"
	"fmt"
	"html"
	"time"

	"github.com/someshubham/gator/data"
	"github.com/someshubham/gator/internal/database"
)

func handlerAgg(s *state, _ command) error {
	ticker := time.NewTicker(time.Duration(60 * time.Second))
	for ; ; <-ticker.C {
		scrapeFeeds(*s.db)
	}
}

func purifyFeed(rssFeed data.RSSFeed) data.RSSFeed {

	var feed data.RSSFeed

	feed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	feed.Channel.Link = html.UnescapeString(rssFeed.Channel.Link)
	feed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for _, item := range rssFeed.Channel.Item {
		var newItem data.RSSItem
		newItem.Description = html.UnescapeString(item.Description)
		newItem.Link = html.UnescapeString(item.Link)
		newItem.Title = html.UnescapeString(item.Title)

		feed.Channel.Item = append(feed.Channel.Item, newItem)
	}

	return feed
}

func scrapeFeeds(db database.Queries) error {

	nextFeed, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        nextFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("Unable to mark feed as fetched")
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)

	if err != nil {
		return err
	}

	purifiedFeed := purifyFeed(*rssFeed)

	for _, item := range purifiedFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
