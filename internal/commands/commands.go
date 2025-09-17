package commands

import (
	"errors"

	"gitlab.com/MadhurSahu/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	commands map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, exists := c.commands[cmd.Name]

	if !exists {
		return errors.New("command not found")
	}

	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.commands[name] = f
}

func NewCommandsInstance() Commands {
	commands := Commands{
		commands: make(map[string]func(*state.State, Command) error),
	}

	commands.Register("register", handlerRegister)
	commands.Register("login", handlerLogin)
	commands.Register("reset", handlerReset)
	commands.Register("users", handlerListUsers)
	commands.Register("agg", handlerFetchFeeds)
	commands.Register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.Register("feeds", handlerListFeeds)
	commands.Register("follow", middlewareLoggedIn(handlerFollowFeed))
	commands.Register("following", middlewareLoggedIn(handlerListUserFeedFollows))
	commands.Register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	commands.Register("browse", middlewareLoggedIn(handlerBrowseFeeds))

	return commands
}
