package main

import (
	"fmt"
	"os"

	"github.com/pderyuga/pokedex-go/internal/pokecache"
)

func commandExit(param string, config *Config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
