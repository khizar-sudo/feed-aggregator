package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/khizar-sudo/feed-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	var limit int32
	if len(cmd.args) > 0 {
		num, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(num)
	} else {
		limit = 2
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, val := range posts {
		fmt.Printf("* %s\n", val.Title.String)
	}
	return nil
}
