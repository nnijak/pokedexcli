package main

import (
	"testing"
	"time"

	"github.com/nnijak/pokedexcli/internal/pokecache"
)

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area/", []byte("testdata"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area/")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
