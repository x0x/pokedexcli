package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/pokeapi"

)

func StartRepl(pokeapiClient *pokeapi.Client) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedexcli> ")
		scanner.Scan()
		text := scanner.Text()

		commands, exists := getCommands()[text]

		if exists {
			err := commands.callBack(pokeapiClient)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callBack    func(*pokeapi.Client) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callBack:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callBack:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callBack:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callBack:    commandMapB,
		},
	}
}
