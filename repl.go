package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	commands := getCommands()
	c := Config{}

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
				err := command.callback(&c)
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
	callback    func(c *Config) error
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
	}
}
