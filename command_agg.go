package main

import (
	"context"
	"database/sql"
	"fmt"
	"html"
	"time"

	"github.com/google/uuid"
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
		fmt.Println(item.PubDate)
		db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			FeedID:      nextFeed.ID,
			Title:       getSqlString(item.Title),
			Description: getSqlString(item.Description),
			Url:         item.Link,
			PublishedAt: getSqlTime(item.PubDate),
		})
	}

	return nil
}

func getSqlString(item string) sql.NullString {
	str := sql.NullString{Valid: false}

	if len(item) != 0 {
		str = sql.NullString{
			String: item,
			Valid:  true,
		}
	}

	return str
}

func getSqlTime(timeStr string) sql.NullTime {
	layout := "2026-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return sql.NullTime{Valid: false}
	}

	// 2. Wrap into sql.NullTime
	nullTime := sql.NullTime{
		Time:  parsedTime,
		Valid: true, // Tells the database driver this is NOT a NULL value
	}

	return nullTime
}
