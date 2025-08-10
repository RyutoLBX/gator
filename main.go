package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg := config.Read()

	cfg.SetUser("ryuto")

	cfg = config.Read()
	fmt.Printf("%+v", cfg)
}
