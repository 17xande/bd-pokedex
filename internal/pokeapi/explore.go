package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Explore(location string) (RespEncounter, error) {
	url := baseURL + "/location-area/" + location

	dat, ok := c.cache.Get(url)
	if !ok {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespEncounter{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespEncounter{}, err
		}

		defer resp.Body.Close()
		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespEncounter{}, err
		}

		c.cache.Add(url, dat)
	}

	locationsResp := RespEncounter{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespEncounter{}, err
	}

	return locationsResp, nil
}
