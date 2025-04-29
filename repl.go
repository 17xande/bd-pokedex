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
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		text := scanner.Text()
		clean := cleanInput(text)
		if len(clean) == 0 {
			continue
		}

		command, ok := getCommands()[clean[0]]
		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}

		if err := command.callback(&conf); err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Displays the names of the next location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous location areas",
			callback:    commandMapB,
		},
	}
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
