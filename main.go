package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/khizar-sudo/feed-aggregator/internal/config"
	"github.com/khizar-sudo/feed-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	// initialise json config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//initialise database and state (pointer to config and database)
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	// initialise the map of commands
	commands := commands{
		c: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

	// fetch CLI arguments
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Insufficient arguments!")
	}

	// prepare to execute command by parsing arguments
	cmd := command{
		name: args[1],
		args: args[2:],
	}

	// run the command
	err = commands.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
