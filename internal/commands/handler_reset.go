package commands

import (
	"context"
	"fmt"

	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerReset(s *state.State, cmd Command) error {
	if err := s.DB.PurgeUsers(context.Background()); err != nil {
		return err
	}

	fmt.Println("Users purged")
	return nil
}
