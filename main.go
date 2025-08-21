package main

import (
	"database/sql"
	"gator/internal/commands"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Error: not enough arguments")
	}

	cfg := config.Read()
	dbURL := cfg.GetDatabaseURL()
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	dbQueries := database.New(db)
	state := commands.NewState(&cfg, dbQueries)

	cmds := commands.NewCommands()
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerUsers)

	name := os.Args[1]
	args := os.Args[2:]
	command := commands.NewCommand(name, args)
	err = cmds.Run(&state, command)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
