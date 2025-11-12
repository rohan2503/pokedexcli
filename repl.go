package main

import (
	"strings"
	"unicode"
)

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, text)
	//remove all the non-alphanumeric characters

	if cleaned == "" {
		return []string{}
	}
	return strings.Fields(cleaned)
}
