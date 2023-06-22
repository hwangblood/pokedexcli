package pokeapi

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"` // sometimes be nil, when json value is null
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
