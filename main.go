package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedox > ")

		if !scanner.Scan() {
			fmt.Println("\nExiting pokedox...")
			break
		}

		text := scanner.Text()

		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s\n", words[0])

	}

}
