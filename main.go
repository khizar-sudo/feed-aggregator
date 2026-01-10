package main

import (
	"fmt"
	"log"
	"os"

	"github.com/khizar-sudo/feed-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
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
		fmt.Println("Insufficient arguments!")
		os.Exit(1)
	}

	// prepare to execute command by parsing arguments
	cmd := command{
		name: args[1],
	}
	if len(args) > 2 {
		cmd.args = args[2:]
	} else {
		cmd.args = nil
	}

	// run the command
	err = commands.run(&s, cmd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Need a username to login!")
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User successfully set to %s\n", cmd.args[0])

	return nil
}
