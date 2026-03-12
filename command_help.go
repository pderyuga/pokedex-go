package main

import (
	"fmt"

	"github.com/pderyuga/pokedex-go/internal/pokecache"
)

func commandHelp(config *Config, cache *pokecache.Cache) error {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}
