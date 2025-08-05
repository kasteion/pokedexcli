package main

import (
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	text = strings.ToLower(text)
	return strings.Split(text, " ")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name: "explore",
			description: "Displays pokemon encounters in a location area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catches a pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a pokemon in your pokedex",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "List pokemons caught",
			callback: commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
