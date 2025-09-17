package commands

import (
	"context"
	"fmt"

	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerListUsers(s *state.State, cmd Command) error {
	users, err := s.DB.ListUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
