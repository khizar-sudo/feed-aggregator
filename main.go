package main

import (
	"log"
	"os"

	"github.com/khizar-sudo/feed-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	// initialise json config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//initialise state (just a pointer to the config)
	s := state{
		cfg: &cfg,
	}

	// initialise the map of commands
	commands := commands{
		c: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

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
