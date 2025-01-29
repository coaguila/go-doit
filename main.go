package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-doit/utils"
)

// Error handling
var errMsg = map[int]string{
	utils.NoErr:           "",
	utils.NotValidCommand: " NotValidCommand Error: Command is not valid",
	utils.NotValidArgs:    "NotValidArgs: Arguments passed are not valid",
}

// Handling commands
var commands = map[string]func([]string){ //TODO: add functions here eventually mdr

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" Welcome to go-doit!")
	fmt.Println(" Choose what to do: (1: Create Todo) (2: Check Todo)")
	// Generate loop for functions
	for {

		// Receive and clean new input
		fmt.Print(" > ")
		raw, _ := reader.ReadString('\n')
		rawN := strings.TrimRight(raw, "\n")
		args := strings.Split(rawN, " ")

		//  Check quit
		if args[0] == "quit" {
			break
		}

		// Check for echo
		if args[0] == "echo" {
			fmt.Print(" ")
			fmt.Println(rawN[5:])
			continue
		}

		// Grab functions
		com, comExists := commands[args[0]]

		if !comExists {
			fmt.Println(errMsg[utils.NotValidCommand])
			continue
		}

		com(args[1:])
	}

}
