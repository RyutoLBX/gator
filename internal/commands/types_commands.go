package commands

import "gator/internal/config"

type state struct {
	config *config.Config
}

func NewState(c config.Config) state {
	return state{config: &c}
}

type command struct {
	name string
	args []string
}

func NewCommand(name string, args []string) command {
	return command{name: name, args: args}
}

type commands struct {
	CommandMap map[string]func(*state, command) error
}

func NewCommands() commands {
	return commands{map[string]func(*state, command) error{}}
}
