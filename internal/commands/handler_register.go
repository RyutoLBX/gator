package commands

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func HandlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("a name is required")
	}

	name := cmd.args[0]

	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	dbUser, err := s.db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return err
	}

	err = s.config.SetUser(dbUser.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User \"%s\" has been created!\n", name)
	fmt.Println("User data:")
	fmt.Println("- ID        :", dbUser.ID)
	fmt.Println("- created_at:", dbUser.CreatedAt)
	fmt.Println("- updated_at:", dbUser.UpdatedAt)
	fmt.Println("- Name      :", dbUser.Name)
	return nil
}
