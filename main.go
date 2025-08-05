package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kasteion/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)

		command, exist := getCommands()[input[0]]
		if exist {
			switch command.name {
			case "explore":
				pokeapi.GetPokeAPIClient().LocationAreaName = input[1]
			case "catch":
				fallthrough	
			case "inspect":
				pokeapi.GetPokeAPIClient().PokemonName = input[1]
			}


			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
