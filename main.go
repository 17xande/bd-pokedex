package main

import "strings"

func main() {
	println("Hello, World!")
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
