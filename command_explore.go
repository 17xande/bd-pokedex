package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location provided")
	}

	location := args[0]
	exploreRes, err := cfg.pokeapiClient.Explore(location)
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
