package main

import "fmt"

func commandHelp(conf *config) error {
	text := `Welcome to the Pokedex!
Usage:

`
	for _, c := range getCommands() {
		text += fmt.Sprintf("%s: %s\n", c.name, c.description)
	}

	fmt.Print(text)
	return nil
}
