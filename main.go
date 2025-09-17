package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gitlab.com/MadhurSahu/gator/internal/database"
)
import (
	"log"
	"os"

	"gitlab.com/MadhurSahu/gator/internal/commands"
	"gitlab.com/MadhurSahu/gator/internal/config"
	"gitlab.com/MadhurSahu/gator/internal/state"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No command provided")
	}

	appConfig, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", appConfig.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	appState := &state.State{
		Config: &appConfig,
		DB:     database.New(db),
	}

	appCommands := commands.NewCommandsInstance()

	command := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := appCommands.Run(appState, command); err != nil {
		log.Fatal(err)
	}
}
