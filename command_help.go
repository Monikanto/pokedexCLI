package main

import "fmt"

func callbackHelp() {

	fmt.Println("welcome to the prokedex help menu:")
	fmt.Println("here are yours available commands:")

	availableCommands := getCommand()

	for _, cmd := range availableCommands{
		fmt.Printf(" - %s: %s", cmd.name , cmd.description)
	}

	
	fmt.Println("")
}

