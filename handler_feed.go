package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Kwasos/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
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

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please speciy a single URL")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}
	row, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Printf("user: %s\n", row.UserName)
	fmt.Printf("feed: %s\n", row.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, currentUser database.User) error {
	rows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.Name)
	if err != nil {
		return fmt.Errorf("error getting follows: %v", err)
	}
	if len(rows) == 0 {
		fmt.Printf("%s is not following any feeds\n", currentUser.Name)
		return nil
	}
	fmt.Printf("Feeds followed by %s:\n", currentUser.Name)
	for _, row := range rows {
		fmt.Printf("* %s\n", row.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please speciy a feed")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}
	p := database.UnfollowParams{
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}
	err = s.db.Unfollow(context.Background(), p)
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}
	return nil
}
