package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// fetch location areas with the reqURL, if nil, fetch the the first page of location areas
// https://pokeapi.co/docs/v2#location-areas
func (c *Client) ListLocationAreas(reqURL *string) (LocationAreasResp, error) {
	endpoint := "location-area"
	fullURL, err := url.JoinPath(baseURL, endpoint)
	if err != nil {
		return LocationAreasResp{}, err
	}
	if reqURL != nil {
		fullURL = *reqURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
