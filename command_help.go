package main

import "fmt"

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}
