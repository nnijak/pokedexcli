package main

import (
	"fmt"
)

func commandHelp(cfg *config, parameter string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, command := range supportedCommands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}
