package pokeapi

import (
	"net/http"
	"time"

	"github.com/kasteion/pokedexcli/internal/pokecache"
)

type pokeAPI struct {
	httpClient http.Client
	BaseURL string
	LocationAreaPage int
	LocationAreaName string
	PokemonName string
	cache pokecache.Cache
	Pokedex Pokedex
}

var client *pokeAPI

func GetPokeAPIClient() *pokeAPI {
	if client == nil {
		client = &pokeAPI{ 
			httpClient: http.Client{
				Timeout: 5 * time.Second,
			}, 
			BaseURL: pokeApiBaseURL, 
			LocationAreaPage: -1,
			cache: pokecache.NewCache(5 * time.Minute),
			Pokedex: newPokedex(),
		}
	}
	return client
}
