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

func handlerFollowFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("follow command requires a feed URL to follow")
	}

	feed, err := s.DB.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	feedFollow, err := s.DB.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	fmt.Println(feedFollow.UserName)
	fmt.Println(feedFollow.FeedName)

	return nil
}
