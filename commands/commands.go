package commands

import (
	"database/sql"
	"fmt"
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

type command struct {
	name string
	args []string
}

type commands struct {
	c map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	val, ok := c.c[cmd.name]
	if !ok {
		return fmt.Errorf("Command does not exist")
	}

	err := val(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.c[name] = f
}

func Init() {
	// initialise json config (db url and current user name)
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
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)

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
