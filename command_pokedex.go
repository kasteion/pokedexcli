package main

import (
	"fmt"

	"github.com/kasteion/pokedexcli/internal/pokeapi"
)

func commandPokedex() error {
	api := pokeapi.GetPokeAPIClient()

	fmt.Println("Your Pokedex:")

	ch, length := api.Pokedex.GetPokemonsChannel()
	for range length {
		pokemon := <- ch
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return  nil
}