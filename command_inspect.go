package main

import (
	"errors"
	"fmt"

	"github.com/pderyuga/pokedex-go/internal/pokecache"
)

func CommandInspect(param string, config *Config, cache *pokecache.Cache, pokedex Pokedex) error {
	if param == "" {
		return errors.New("you must provide a pokemon name")
	}

	pokemon, ok := pokedex[param]
	if !ok {
		return fmt.Errorf("You have not caught %s", param)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
