package main

import (
	"pokedexcli/pokeapi"
	"time"
)


func main() {
 
  pokeapiClient := pokeapi.NewClient(5*time.Second,5*time.Minute)
	StartRepl(&pokeapiClient)
}

