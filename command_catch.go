package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	chance := 0
	// for balls := 1; balls > 0; balls-- {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	chance = rand.IntN(pokemonResp.BaseExperience / 20)
	if chance == 1 {
		fmt.Printf("%s caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemonResp
		return nil
	}
	fmt.Println("missed")
	// }

	// fmt.Printf("%s got away.\n", pokemonName)

	return nil
}
