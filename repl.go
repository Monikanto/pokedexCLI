package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >  ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0{
			continue
			
		}

		command := cleaned[0]

		switch command {
		case "help":
			fmt.Println("welcome to the prokedex help menu:")
			fmt.Println("here are yours available commands:")
			fmt.Println(" -help")
			fmt.Println(" -exit")
			fmt.Println("")

		case "exit":
			os.Exit(0) //terminates the program

		default:
			fmt.Println("invaild commmand")
		}

		
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}