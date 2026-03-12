package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandExit,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

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
			}

			switch command.name {
			case "help":
				commandHelp(commands)
			case "exit":
				commandExit()
			}
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	parts := strings.Fields(lowercaseText)
	return parts
}
