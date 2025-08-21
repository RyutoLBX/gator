package commands

import (
	"context"
	"fmt"
)

func HandlerReset(s *state, cmd command) error {
	s.db.DeleteAllUsers(context.Background())
	fmt.Println("Deleted all entries from users table.")
	return nil
}
