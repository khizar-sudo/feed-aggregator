package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/khizar-sudo/feed-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Need a URL to follow!")
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	
	fmt.Printf("* Feed Name: %s\n", feedFollow.FeedName)
	fmt.Printf("* User Name: %s\n", feedFollow.UserName)
	fmt.Println("=====================================")

	return nil
}