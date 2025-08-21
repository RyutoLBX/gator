package commands

import (
	"gator/internal/config"
	"gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func NewState(c *config.Config, db *database.Queries) state {
	return state{config: c, db: db}
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
