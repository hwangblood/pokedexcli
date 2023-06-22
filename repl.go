package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type commandCallback = func() error

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
	}
}

func startRepl() {
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
		command.callback()
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
