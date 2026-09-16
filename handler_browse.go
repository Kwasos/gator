package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Kwasos/gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	userName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	limit := 2
	if len(cmd.Args) == 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
	}
	posts, err := s.db.GetPostsByUser(context.Background(), database.GetPostsByUserParams{
		UserID: user.ID, Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting user posts: %v", err)
	}
	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\n", post.Title, post.Url)
	}
	return nil
}
