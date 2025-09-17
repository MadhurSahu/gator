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

func handlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register command requires a username")
	}

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	if err := s.Config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User %s is registered\n", user.Name)
	return nil
}
