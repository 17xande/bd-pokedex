package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location := args[0]
	exploreRes, err := cfg.pokeapiClient.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", exploreRes.Location.Name)
	fmt.Printf("Found Pokemon:\n")
	for _, e := range exploreRes.PokemonEncounters {
		fmt.Printf(" - %s\n", e.Pokemon.Name)
	}
	return nil
}
