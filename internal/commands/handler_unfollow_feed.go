package commands

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/MadhurSahu/gator/internal/database"
	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerUnfollowFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("unfollow command requires a feed URL to unfollow")
	}

	feed, err := s.DB.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	err = s.DB.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s unfollowed %s\n", user.Name, feed.Name)
	return nil
}
