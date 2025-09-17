package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.com/MadhurSahu/gator/internal/database"
	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerAddFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed command requires a name and a feed URL")
	}

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	_, err = s.DB.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
