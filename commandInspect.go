package main

import (
	"fmt"

	"pokedexcli/pokeapi"
)

func commandInspect(pokeapiClient *pokeapi.Client, args ...string) error {
	data, ok := userPokedex[args[0]]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name: ", data.Name)
	fmt.Println("Height: ", data.Height)
	fmt.Println("Weight: ", data.Weight)
	fmt.Println("Stats: ")

	for _, v := range data.Stats {
		fmt.Println("\t", v.Stat.Name, ": ", v.BaseStat)
	}

	fmt.Println("Types: ")
	for _, v := range data.Types {
		fmt.Println("\t", v.Type.Name)
	}

	return nil
}
