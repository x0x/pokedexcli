package main

import (
	"fmt"
	"math/rand"

	"pokedexcli/pokeapi"
)

var userPokedex map[string]pokeapi.CatchResponse = make(map[string]pokeapi.CatchResponse)

func commandCatch(pokeapiClient *pokeapi.Client, args ...string) error {
	fmt.Println("Throwing a Pokeball at", args[0], "...")
	response, err := pokeapi.CatchPokemon(pokeapiClient, &args[0])
	if err != nil {
		return err
	}

	baseExperience := response.BaseExperience
	randomNumber := rand.Intn(baseExperience)
	if randomNumber >= 40 {
		fmt.Println(args[0], " was caught!")
		userPokedex[args[0]] = response
	} else {
		fmt.Println(args[0], " escaped!")
	}

	return nil
}
