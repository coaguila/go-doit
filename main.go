package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command int

const (
	Echo Command = 0
	CreateTodo
	Quit
)

type Todo struct {
	Id    string
	User  string
	Title string
	Task  string
	Done  bool
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to go-doit!")

	for {
		fmt.Println("Choose what to do: (1: Create Todo) (2: Check Todo)")
		input, _ := reader.ReadString('\n')
		req := strings.TrimRight(input, "\n")

		if req == "q" {
			break
		}
		fmt.Println(req)
	}

}
