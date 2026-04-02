package main

import (
	"time"

	"github.com/nnijak/pokedexcli/internal/pokecache"
)

var supportedCommands map[string]cliCommand
var newCache *pokecache.Cache
var pokeBag map[string]Pokemon

func main() {

	supportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous page of location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemons in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw a pokeball",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show pokemon details",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all caught pokemon",
			callback:    commandPokedex,
		},
	}

	newCache = pokecache.NewCache(60 * time.Second)
	pokeBag = make(map[string]Pokemon)
	replStart()
}
