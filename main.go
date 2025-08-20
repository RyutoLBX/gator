package main

import (
	"gator/internal/commands"
	"gator/internal/config"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Error: not enough arguments")
	}
	cfg := config.Read()

	state := commands.NewState(cfg)

	cmds := commands.NewCommands()
	cmds.Register("login", commands.HandlerLogin)

	name := os.Args[1]
	args := os.Args[2:]
	command := commands.NewCommand(name, args)
	err := cmds.Run(&state, command)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
