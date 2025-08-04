package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p *pokeAPI) GetLocationAreas(page int) ([]LocationArea, error){
	url := fmt.Sprintf("%s%s?offset=%v", p.BaseURL, locationAreaPath, 20 * page)
	
	cacheVal, exist := p.cache.Get(url)
	if exist {
		var body GetLocationAreasResponse
		if err := json.Unmarshal(cacheVal, &body); err != nil {
			return nil, err
		}

		return body.Results, nil
	}

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

	p.cache.Add(url, data)

	var body GetLocationAreasResponse
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}

	return body.Results, nil
}