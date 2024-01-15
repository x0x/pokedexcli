package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedexcli>")
		scanner.Scan()
		text := scanner.Text()

		commands, exists := getCommands()[text]

		if exists {
			err := commands.callBack()
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
	callBack    func() error
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
	}
}
