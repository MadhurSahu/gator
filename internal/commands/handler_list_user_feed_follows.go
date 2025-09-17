package commands

import (
	"context"

	"gitlab.com/MadhurSahu/gator/internal/database"
	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerListUserFeedFollows(s *state.State, cmd Command, user database.User) error {
	feedFollows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	println(user.Name)
	for _, feedFollow := range feedFollows {
		println(feedFollow.FeedName)
	}

	return nil
}
