package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	parts := strings.Fields(lowercaseText)
	return parts
}
