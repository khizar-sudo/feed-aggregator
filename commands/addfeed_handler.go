package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/khizar-sudo/feed-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 || len(cmd.args) > 2 {
		return fmt.Errorf("Insufficient arguments!")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})

	fmt.Printf("Feed fetched %+v\n", feed)

	return nil
}
