package main

import (
	"os"
	"pokedexcli/pokeapi"
)

func commandExit(pokeapiClient *pokeapi.Client, args ...string) error {
	os.Exit(0)
	return nil
}
