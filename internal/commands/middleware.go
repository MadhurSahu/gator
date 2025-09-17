package commands

import (
	"context"

	"gitlab.com/MadhurSahu/gator/internal/database"
	"gitlab.com/MadhurSahu/gator/internal/state"
)

func middlewareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
	return func(s *state.State, cmd Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
