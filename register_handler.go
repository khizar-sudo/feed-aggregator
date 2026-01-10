package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/khizar-sudo/feed-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Need a user to register!")
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("User created successfully: %v\n", user)
	return nil
}
