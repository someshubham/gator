package main

import (
	"context"
	"fmt"
	"html"
	"os"

	"github.com/someshubham/gator/data"
)

func handlerAgg(_ *state, _ command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(purifyFeed(*rssFeed))
	return nil
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
