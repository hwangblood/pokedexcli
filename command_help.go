package main

import "fmt"

func callBackHelp() error {
	fmt.Println("Welcome to the Pokédex help manual!")
	fmt.Println("Here are available commands:")
	availableCmds := getCommands()
	for _, cmd := range availableCmds {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("") // make a empty new line
	return nil
}
