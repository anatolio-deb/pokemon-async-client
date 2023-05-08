package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	ID             int    `json:"id"`
	IsDefault      bool   `json:"is_default"`
	Name           string `json:"name"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemons struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous any      `json:"previous"`
	Results  []Result `json:"results"`
}

func getPokemon(url string) Pokemon {
	var p Pokemon
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &p)

	if err != nil {
		panic(err)
	}

	return p

}

func GetPokemons() []Pokemon {
	var p Pokemons
	var pokemons []Pokemon
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=0&limit=1281")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &p)

	if err != nil {
		panic(err)
	}

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, p := range p.Results {
		wg.Add(1)
		go func(p Result) {
			defer wg.Done()
			pokemon := getPokemon(p.Url)
			mu.Lock()
			pokemons = append(pokemons, pokemon)
			mu.Unlock()
		}(p)
	}
	wg.Wait()

	return pokemons
}
