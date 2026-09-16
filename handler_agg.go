package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Kwasos/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("please specify a time duration (like '1m' for 1 minute)")
	}
	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("couldn't get next feed to fetch: %v", err)
		return
	}
	_, err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed.ID,
	})
	if err != nil {
		log.Printf("couldn't get next feed to fetch: %v", err)
		return
	}
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't fetch feed: %v", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		log.Printf("Fetched: %v", item.Title)
	}
}
