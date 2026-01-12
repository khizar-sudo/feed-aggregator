package commands

import (
	"context"
	"fmt"

	"github.com/khizar-sudo/feed-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Need a username to login!")
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Unfollowed feed successfully: %s\n", feed.Url)
	return nil
}
