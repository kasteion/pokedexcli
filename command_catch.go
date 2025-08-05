package main

import (
	"fmt"
	"math/rand"

	"github.com/kasteion/pokedexcli/internal/pokeapi"
)

func commandCatch() error {
	api := pokeapi.GetPokeAPIClient()
	pokemon, err := api.GetPokemon(api.PokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", api.PokemonName)
	
	luck := rand.Intn(pokemon.BaseExperience)
	if luck > 40 {
		fmt.Printf("%s escaped!\n", api.PokemonName)
		return nil
	}

	api.Pokedex.Add(api.PokemonName, pokemon)
	fmt.Printf("%s was caught!\n", api.PokemonName)
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}