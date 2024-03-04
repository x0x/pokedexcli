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

func commandMap() error {

	fmt.Println("hello world Welcome to Map")

	response, err := pokeapi.GetLocations(cfg.nextLocationUrl)

	if err != nil {
		return nil
	}

	PrettyPrint(response)

	cfg.prevLocationUrl = response.Previous
	cfg.nextLocationUrl = response.Next

	return nil
}

func commandMapB() error {
	fmt.Println("Hello to Mapb")

	response, err := pokeapi.GetLocations(cfg.prevLocationUrl)

	if err != nil {
		return err
	}

	PrettyPrint(response)
  cfg.prevLocationUrl = response.Previous
	cfg.nextLocationUrl = response.Next

	return nil
}

func PrettyPrint(i pokeapi.Response) {
	for _, value := range i.Results {
		fmt.Println(value.Name)
	}
}
