package commands

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("%s command argument cannot be empty", cmd.name)
	}

	username := cmd.args[0]
	s.config.CurrentUserName = username
	fmt.Printf("Current username has been set to %s!", username)
	return nil
}
