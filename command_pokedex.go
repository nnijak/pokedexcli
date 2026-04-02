package main

import "fmt"

func commandPokedex(cfg *config, parameter string) error {
	if len(pokeBag) == 0 {
		fmt.Printf("You haven't caught any pokemon yet!\n")
		return nil
	}
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range pokeBag {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
