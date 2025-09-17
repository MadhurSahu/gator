package commands

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login command requires a username")
	}

	user, err := s.DB.GetUser(context.Background(), cmd.Args[0])

	if err != nil {
		return err
	}

	if err := s.Config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("Logged in as %s\n", user.Name)
	return nil
}
