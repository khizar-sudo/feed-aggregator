package commands

import (
	"context"
	"fmt"

	"github.com/khizar-sudo/feed-aggregator/feed"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
