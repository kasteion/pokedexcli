package main

import (
	"fmt"

	"github.com/kasteion/pokedexcli/internal/pokeapi"
)

func commandInspect() error {
	api := pokeapi.GetPokeAPIClient()
	pokemon, ok := api.Pokedex.Get(api.PokemonName)
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range(pokemon.Stats) {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	
	for _, pType := range(pokemon.Types) {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil 
}