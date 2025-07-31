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
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
    	"exit": {
    	    name:        "exit",
    	    description: "Exit the Pokedex",
   	     callback:    commandExit,
    	},
	}
}
