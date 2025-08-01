package pokeapi

import (
	"net/http"
	"time"
)

type pokeAPI struct {
	httpClient http.Client
	BaseURL string
	LocationAreaPage int
}

var client *pokeAPI

func GetPokeAPIClient() *pokeAPI {
	if client == nil {
		client = &pokeAPI{ httpClient: http.Client{
			Timeout: 5 * time.Second,
		}, BaseURL: pokeApiBaseURL, LocationAreaPage: -1}
	}
	return client
}
