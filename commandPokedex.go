package main

import (
	"fmt"

	"pokedexcli/pokeapi"
)

func commandPokedex(pokeapiClient *pokeapi.Client, args ...string) error {
	fmt.Println("Your Pokedex: ")
	for k := range userPokedex {
		fmt.Println("\t", k)
	}

	return nil
}
