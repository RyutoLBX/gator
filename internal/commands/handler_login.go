package commands

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("a username is required")
	}

	username := cmd.args[0]
	s.config.SetUser(username)
	fmt.Printf("Current username has been set to %s!\n", username)
	return nil
}
