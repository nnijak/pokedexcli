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
	callback    func(*config, string) error
}

type config struct {
	next     string
	previous string
}

var initConfig = config{
	next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	previous: "",
}

func replStart() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		parameter := ""
		text := scanner.Text()
		splittedText := strings.Fields(text)
		typedCommand := splittedText[0]
		if len(splittedText) > 1 {
			parameter = splittedText[1]
		}
		for _, command := range supportedCommands {
			if command.name == typedCommand {
				command.callback(&initConfig, parameter)
			}
		}
	}
}
