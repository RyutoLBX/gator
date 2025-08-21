package commands

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("a username is required")
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	s.config.SetUser(username)
	fmt.Printf("Current username has been set to %s!\n", username)
	return nil
}
