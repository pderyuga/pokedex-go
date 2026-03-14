package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pderyuga/pokedex-go/internal/pokecache"
)

type Pokedex map[string]Pokemon

func newPokedex() Pokedex {
	pokedex := make(map[string]Pokemon)
	return pokedex
}

func startRepl() {
	commands := getCommands()
	pokedex := newPokedex()
	config := Config{}
	cache := pokecache.NewCache(60 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			if len(words) == 0 {
				continue
			}

			firstWord := words[0]
			command, ok := commands[firstWord]
			if !ok {
				fmt.Printf("Unknown command: %s\n", firstWord)
				continue
			} else {
				param := ""
				if len(words) > 1 {
					param = words[1]
				}
				err := command.callback(param, &config, &cache, pokedex)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}

		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	parts := strings.Fields(lowercaseText)
	return parts
}

type cliCommand struct {
	name        string
	description string
	callback    func(param string, config *Config, cache *pokecache.Cache, pokedex Pokedex) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Explore the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore the previous 20 locations the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Explore a specific location in the Pokemon world",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempt to catch a Pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Inspect a caught Pokemon",
			callback:    CommandInspect,
		},
	}
}
