package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p *pokeAPI) GetLocationAreas(page int) ([]LocationArea, error){
	url := fmt.Sprintf("%s%s?offset=%v", p.BaseURL, locationAreaPath, 20 * page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// This works too...
	// var body GetLocationAreaResponse
	// if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
	// 	return nil, err
	// }

	// return body.Results, nil

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var body GetLocationAreaResponse
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}

	return body.Results, nil
}