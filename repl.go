package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type commandCallback = func(*config) error

type cliCommand struct {
	name        string
	description string
	callback    commandCallback
}

// return map , key is the name of the command, value is the command
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help information",
			callback:    callBackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokédex REPL",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous page of location areas",
			callback:    callbackMapBack,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	availableCmds := getCommands()

	for {
		fmt.Print(" > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		// when type ENTER without any input
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		command, ok := availableCmds[commandName]

		if !ok {
			fmt.Printf("invalid command: %v\n", commandName)
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
