package commands

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/khizar-sudo/feed-aggregator/feed"
	"github.com/khizar-sudo/feed-aggregator/internal/database"
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
			publishedAt := sql.NullTime{}
			if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
				publishedAt = sql.NullTime{Time: t, Valid: true}
			} else {
				fmt.Printf("Could not parse date: %v\n", item.PubDate)
			}

			_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       sql.NullString{String: item.Title, Valid: true},
				Url:         item.Link,
				Description: sql.NullString{String: item.Description, Valid: true},
				PublishedAt: publishedAt,
				FeedID:      feedToFetch.ID,
			})
			if err != nil {
				if !strings.Contains(err.Error(), "duplicate key") &&
					!strings.Contains(err.Error(), "unique constraint") {
					fmt.Println(err.Error())
				}
			} else {
				fmt.Printf("Post added: %s\n", item.Title)
			}
		}
		fmt.Println("=====================================")
	}
}
