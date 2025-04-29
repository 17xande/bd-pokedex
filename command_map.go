package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationResponse struct {
	Previous string
	Next     string
	Results  []locationArea
}

type locationArea struct {
	Name string
	URL  string
}

func commandMap(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if conf.next != "" {
		url = conf.next
	}

	return mapGet(conf, url)
}

func commandMapB(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if conf.previous != "" {
		url = conf.previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	return mapGet(conf, url)
}

func mapGet(conf *config, url string) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	locRes := locationResponse{}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locRes)
	if err != nil {
		return err
	}

	conf.next = locRes.Next
	conf.previous = locRes.Previous

	for _, l := range locRes.Results {
		fmt.Println(l.Name)
	}

	return nil

}
