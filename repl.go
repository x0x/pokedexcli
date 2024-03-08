package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/pokeapi"
	"strings"
)

func StartRepl(pokeapiClient *pokeapi.Client) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedexcli> ")
		scanner.Scan()
		text := scanner.Text()
    input := strings.Fields(text)
		commands, exists := getCommands()[input[0]]
    
    args := []string{}
    if len(input) > 1 {
      args = input[1 : ]   
    }

		if exists {
			err := commands.callBack(pokeapiClient, args...)
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
	callBack    func(*pokeapi.Client, ...string) error
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
    "explore": {
 			name:        "explore",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callBack:    commandExplore,
	   },

	}
}
