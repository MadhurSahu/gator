package commands

import (
	"context"
	"fmt"

	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerListFeeds(s *state.State, cmd Command) error {
	feeds, err := s.DB.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserName)
	}

	return nil
}
