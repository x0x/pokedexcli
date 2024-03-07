package main

import (
	"os"
	"pokedexcli/pokeapi"
)

func commandExit(pokeapiClient *pokeapi.Client) error {
	os.Exit(0)
	return nil
}
