package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/someshubham/gator/data"
)

func fetchFeed(ctx context.Context, feedURL string) (*data.RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &data.RSSFeed{}, fmt.Errorf("Unable to create request %w", err)
	}

	client := &http.Client{}

	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		return &data.RSSFeed{}, fmt.Errorf("Unable to fetch response %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return &data.RSSFeed{}, fmt.Errorf("status code error: %s", res.Status)
	}

	val, err := io.ReadAll(res.Body)

	if err != nil {
		return &data.RSSFeed{}, fmt.Errorf("Unable to read response %w", err)
	}

	var rssFeed data.RSSFeed

	if err := xml.Unmarshal(val, &rssFeed); err != nil {
		return &data.RSSFeed{}, fmt.Errorf("Unable to parse response %w", err)
	}

	return &rssFeed, nil
}
