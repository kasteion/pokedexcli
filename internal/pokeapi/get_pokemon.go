package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p *pokeAPI) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("%s%s/%s", pokeApiBaseURL, pokemonPath, name)

	cache, ok := p.cache.Get(url)
	if ok {
		var pokemon Pokemon
		if err := json.Unmarshal(cache, &pokemon); err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	p.cache.Add(url, data)

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}