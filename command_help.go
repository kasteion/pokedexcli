package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for k, v := range getCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}

	return nil
}
