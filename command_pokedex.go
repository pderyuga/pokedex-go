package main

import (
	"fmt"

	"github.com/pderyuga/pokedex-go/internal/pokecache"
)

func commandPokedex(param string, config *Config, cache *pokecache.Cache, pokedex Pokedex) error {
	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
