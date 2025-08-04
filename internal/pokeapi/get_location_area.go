package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p *pokeAPI) GetLocationArea(locationID string) (GetLocationAreaResponse, error) {
	url := fmt.Sprintf("%s%s/%s", p.BaseURL, locationAreaPath, locationID)

	cache, ok := p.cache.Get(url)
	if ok {
		var body GetLocationAreaResponse
		if err := json.Unmarshal(cache, &body); err != nil {
			return GetLocationAreaResponse{}, err
		}

		return body, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GetLocationAreaResponse{}, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return GetLocationAreaResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return  GetLocationAreaResponse{}, err
	}

	p.cache.Add(url, data)

	var body GetLocationAreaResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return GetLocationAreaResponse{}, err
	}

	return body, nil
}