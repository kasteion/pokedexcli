package main

import (
	"fmt"

	"github.com/kasteion/pokedexcli/internal/pokeapi"
)

func commandExplore() error {
	api := pokeapi.GetPokeAPIClient()
	locationArea, err := api.GetLocationArea(api.LocationAreaName)
	if err != nil {
		return  err
	}

	fmt.Printf("Exploring %s...\n", api.LocationAreaName)
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounters := range(locationArea.PokemonEncounters) {
		fmt.Printf("- %s\n", pokemonEncounters.Pokemon.Name)
	}

	return nil
}