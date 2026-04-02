package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchLocationAreas(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to get location areas %w", err)
	}
	defer res.Body.Close()

	var locationAreas LocationArea
	if err := json.NewDecoder(res.Body).Decode(&locationAreas); err != nil {
		return LocationArea{}, fmt.Errorf("failed to decode loction areas %w", err)
	}

	jsonData, err := json.Marshal(locationAreas)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to save data to cache %w", err)
	}
	newCache.Add(url, jsonData)

	return locationAreas, nil
}

func commandMap(cfg *config, parameter string) error {
	var locationAreas LocationArea
	locationData, exists := newCache.Get(cfg.next)
	if exists {
		fmt.Println("!LOADED FROM CACHE!")
		if err := json.Unmarshal(locationData, &locationAreas); err != nil {
			return fmt.Errorf("failed to unmarshal location data from cache %w", err)
		}
	} else {
		var err error
		locationAreas, err = fetchLocationAreas(cfg.next)
		if err != nil {
			return fmt.Errorf("failed to fetch location areas: %w", err)
		}
	}
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous
	return nil
}

func commandMapB(cfg *config, parameter string) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	var locationAreas LocationArea
	locationData, exists := newCache.Get(cfg.previous)
	if exists {
		fmt.Println("!LOADED FROM CACHE!")
		if err := json.Unmarshal(locationData, &locationAreas); err != nil {
			return fmt.Errorf("failed to unmarshal location data from cache %w", err)
		}
	} else {
		var err error
		locationAreas, err = fetchLocationAreas(cfg.previous)
		if err != nil {
			return fmt.Errorf("failed to fetch location areas: %w", err)
		}
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous
	return nil
}
