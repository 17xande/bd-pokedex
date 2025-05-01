package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	dat, ok := c.cache.Get(url)
	if !ok {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}

		defer resp.Body.Close()
		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return Location{}, err
		}

		c.cache.Add(url, dat)
	}

	locationsResp := Location{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	return locationsResp, nil
}
