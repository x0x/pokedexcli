package main

import (
	"fmt"
	"pokedexcli/pokeapi"
)

func commandHelp(pokeapiClient *pokeapi.Client) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

