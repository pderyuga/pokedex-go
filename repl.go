package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			if len(words) == 0 {
				continue
			}

			fmt.Printf("Your command was: %s\n", words[0])
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	parts := strings.Fields(lowercaseText)
	return parts
}
