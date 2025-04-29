package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	fmt.Printf("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		clean := cleanInput(text)
		if len(clean) == 0 {
			continue
		}

		command, ok := commands[clean[0]]
		if !ok {
			fmt.Printf("Unknown command\n")
			fmt.Printf("Pokedex > ")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Printf("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func commandHelp() error {
	text := `Welcome to the Pokedex!
Usage:

`
	for _, c := range commands {
		text += fmt.Sprintf("%s: %s\n", c.name, c.description)
	}

	fmt.Print(text)
	return nil
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	strs := strings.Split(text, " ")
	res := []string{}

	for _, s := range strs {
		if s == "" {
			continue
		}
		s = strings.ToLower(s)
		res = append(res, strings.TrimSpace(s))
	}

	return res
}
