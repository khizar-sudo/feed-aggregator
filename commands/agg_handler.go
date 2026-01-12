package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/khizar-sudo/feed-aggregator/feed"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Provide a time duration!")
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	timeDuration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeDuration)
	ticker := time.NewTicker(timeDuration)
	for ; ; <-ticker.C {

		feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
		if err != nil {
			return err
		}

		_, err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
		if err != nil {
			return err
		}

		f, err := feed.FetchFeed(context.Background(), feedToFetch.Url)
		if err != nil {
			return err
		}

		fmt.Printf("Channel: %s\n", f.Channel.Title)
		for _, item := range f.Channel.Item {
			fmt.Printf("* %s\n", item.Title)
		}
		fmt.Println("=====================================")
	}
}
