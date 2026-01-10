package commands

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Need a username to login!")
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User successfully set to %s\n", cmd.args[0])

	return nil
}
