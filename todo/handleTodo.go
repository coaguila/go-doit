package todo

import (
	"fmt"

	. "go-doit/utils"
)

func initTodoEnv() (*SimpleTodoBoard, int) {
	return newSimpleTodoBoard()
}

/*
	Handlers for the commands
*/

func handleShow(args []string, board *SimpleTodoBoard) int {
	if len(args) != 0 {
		return NotValidArgs
	}
	fmt.Println(boardToString(board))
	return NoErr
}

func handleCreateTodo(args []string, board *SimpleTodoBoard) int {
	if len(args) != 2 {
		return NotValidArgs
	}
	return addNewToBoard(board, args[0], args[1])
}
