package main

import "fmt"

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
