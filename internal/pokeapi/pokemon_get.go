package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	pokemon := RespPokemon{}
	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return RespPokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}

	defer resp.Body.Close()
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemon, nil

}
