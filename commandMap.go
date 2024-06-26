package main

import (
	"fmt"

	"pokedexcli/pokeapi"
)

type Config struct {
	nextLocationUrl *string
	prevLocationUrl *string
}

var cfg *Config = &Config{}

func commandMap(pokeapiClient *pokeapi.Client, args ...string) error {
	response, err := pokeapi.GetLocations(pokeapiClient, cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	PrettyPrint(response)

	cfg.prevLocationUrl = response.Previous
	cfg.nextLocationUrl = response.Next

	return nil
}

func commandMapB(pokeapiClient *pokeapi.Client, args ...string) error {
	if cfg.prevLocationUrl == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	response, err := pokeapi.GetLocations(pokeapiClient, cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	PrettyPrint(response)
	cfg.prevLocationUrl = response.Previous
	cfg.nextLocationUrl = response.Next

	return nil
}

func PrettyPrint(i pokeapi.GetResponse) {
	for _, value := range i.Results {
		fmt.Println(value.Name)
	}
}
