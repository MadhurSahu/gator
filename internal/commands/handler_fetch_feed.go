package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gitlab.com/MadhurSahu/gator/internal/database"
	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerFetchFeeds(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("aggregate command requires an interval")
	}

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %dm%ds\n", int(interval.Minutes()), int(interval.Seconds())%60)

	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		feedToFetch, err := s.DB.GetNextFeedToFetch(context.Background())
		if err != nil {
			return err
		}

		feed, err := fetchFeed(context.Background(), feedToFetch.Url)
		if err != nil {
			return err
		}

		println(feed.Channel.Title)
		for _, item := range feed.Channel.Item {
			if item.Title == "" && item.Description == "" {
				continue
			}

			_, err := s.DB.CreatePost(context.Background(), database.CreatePostParams{
				ID:          uuid.New(),
				FeedID:      feedToFetch.ID,
				Title:       item.Title,
				Url:         item.Link,
				Description: item.Description,
				PublishedAt: parseDate(item.PubDate),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})

			if err != nil {
				var err *pq.Error
				fmt.Printf("%s... Failed\n", item.Title)
				if errors.As(err, &err) {
					if err.Code != "23505" {
						fmt.Println(err)
					}
				} else {
					return err
				}
			}

			fmt.Printf("%s... Saved\n", item.Title)
		}

		if err := s.DB.MarkFeedFetched(context.Background(), feedToFetch.ID); err != nil {
			return err
		}
	}
}

func parseDate(date string) sql.NullTime {
	t, err := time.Parse(time.RFC1123Z, date)
	if err != nil {
		return sql.NullTime{Valid: false}
	}

	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
