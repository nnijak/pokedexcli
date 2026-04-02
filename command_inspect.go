package main

import "fmt"

func commandInspect(cfg *config, parameter string) error {
	pokemon, exists := pokeBag[parameter]
	if exists {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokeType.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon")
	}

	return nil
}
