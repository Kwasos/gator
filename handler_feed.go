package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Kwasos/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
	}

	fmt.Printf("Feed created: %+v\n", feed)
	fmt.Printf("Feed followed: %+v\n", follow.FeedName)
	fmt.Printf("Following user: %+v\n", follow.UserName)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	rows, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}
	for _, feed := range rows {
		fmt.Printf("* Name: %s\n", feed.Name)
		fmt.Printf("* URL: %s\n", feed.Url)
		fmt.Printf("* User: %s\n", feed.UserName)
	}
	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("please speciy a single URL")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}
	row, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Printf("user: %s\n", row.UserName)
	fmt.Printf("feed: %s\n", row.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	rows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting follows: %v", err)
	}
	for _, row := range rows {
		fmt.Printf("* Name: %s\n", row.FeedName)
		fmt.Printf("* User: %s\n", row.UserName)
	}
	return nil
}
