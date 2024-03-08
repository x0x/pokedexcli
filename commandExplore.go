package main

import (
	"fmt"

	"pokedexcli/pokeapi"
)

func commandExplore(pokeapiClient *pokeapi.Client, args ...string) error {

  fmt.Println("Exploring",args[0])
  fmt.Println("Found Pokemon ... ")
	response, err := pokeapi.ListLocations(pokeapiClient, args[0])
	if err != nil {
    fmt.Println("Panic error at commandExplore")
		return nil
	}

  PrettyPrintResponse(response)
	return nil
}

func PrettyPrintResponse(i pokeapi.ListResponse) {
	for _, value := range i.PokemonEncounters {
		fmt.Println("-",value.Pokemon.Name)
	}
}
