package todo

import (
	. "go-doit/utils"
)

// This file only handles the concepts of the todo's, CLI calls will be handled somewhere else

/*
	Define constants + types + data structures, current mvp
	TODO: Make more complex structs/types later
*/

type SimpleTodo struct {
	Title string // Title, should not be empty
	Task  string // Descriptions of todo, can be empty
	Done  int    // Should be init to 0 and should not depass 1
}

type SimpleTodoBoard []SimpleTodo

/*
	Define builders for structs
*/

// Creates a new instance of todo
func newSimpleTodo(title string, task string) (*SimpleTodo, int) {
	if title == "" {
		return nil, NotValidParameters
	}

	return &SimpleTodo{title, task, 0}, NoErr
}

// Creates new instance of a simple todo board
func newSimpleTodoBoard() (*SimpleTodoBoard, int) {
	return &SimpleTodoBoard{}, NoErr
}

/*
	related functions
*/

func addNewToBoard(board *SimpleTodoBoard, title string, task string) int {
	newSimpleTodo, err := newSimpleTodo(title, task)
	if err != NoErr {
		return err
	}

	*board = append(*board, *newSimpleTodo)
	return err
}
