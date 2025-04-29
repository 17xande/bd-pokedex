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

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

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

		if err := command.callback(); err != nil {
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
			description: "Displays the names of location areas",
			callback:    commandMap,
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
