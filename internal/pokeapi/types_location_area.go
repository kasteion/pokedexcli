package pokeapi

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`	
}

type GetLocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []LocationArea `json:"results"`
}