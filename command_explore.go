package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Encounters struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, parameter string) error {
	if parameter == "" {
		fmt.Println("Please provide an area you want to explore!")
		return nil
	}

	fmt.Println("Exploring " + parameter + "...")

	fullPath := "https://pokeapi.co/api/v2/location-area/" + parameter
	var encounters Encounters

	exploreData, exists := newCache.Get(fullPath)
	if exists {
		fmt.Println("!LOADED FROM CACHE!")
		if err := json.Unmarshal(exploreData, &encounters); err != nil {
			return fmt.Errorf("failed to unmarshal data for location area: %s, %w", parameter, err)
		}
	} else {
		res, err := http.Get(fullPath)
		if err != nil {
			return fmt.Errorf("failed to fetch data for location area: %s, %w", parameter, err)
		}

		if err := json.NewDecoder(res.Body).Decode(&encounters); err != nil {
			return fmt.Errorf("failed to decode json data for location area: %s, %w", parameter, err)
		}

		jsonData, err := json.Marshal(encounters)
		if err != nil {
			return fmt.Errorf("failed to marshal data for location area: %s, %w", parameter, err)
		}
		newCache.Add(fullPath, jsonData)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}
