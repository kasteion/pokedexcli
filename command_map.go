package main

import (
	"fmt"

	"github.com/kasteion/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	api := pokeapi.GetPokeAPIClient()
	api.LocationAreaPage += 1

	locationAreas, err := api.GetLocationAreas(api.LocationAreaPage)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapB() error {
	api := pokeapi.GetPokeAPIClient()
	if api.LocationAreaPage > 0 {
		api.LocationAreaPage = api.LocationAreaPage - 1
	}

	locationAreas, err := api.GetLocationAreas(api.LocationAreaPage)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	return nil
}
