package state

import (
	"gitlab.com/MadhurSahu/gator/internal/config"
	"gitlab.com/MadhurSahu/gator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}
