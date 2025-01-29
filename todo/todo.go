package todo

import (
	. "go-doit/utils"
)

const (
	COMPLEX_VERSION  string = "0.0.0"
	MAX_BLOCK_SIZE   int    = 20
	MAX_COMPLEX_SIZE int    = 20
)

type SimpleTodo struct {
	Title string // Title, should not be empty
	Task  string // Descriptions of todo, can be empty
	Done  int    // Should be init to 0 and should not depass 1
}

type ComplexMetadata struct {
	Title   string // Title, should not be empty
	Task    string // Description of todo
	Version string // Shows version of complex
	Size    int    // How many sub-SimpleTodo
	Done    int    // Quantity of Sub-SimpleTodo that are done
}

type BlockTodo struct {
	Tasks [MAX_BLOCK_SIZE]SimpleTodo
	Size  int
}

type ComplexTodo struct {
	Metadata ComplexMetadata
	Tasks    [MAX_COMPLEX_SIZE]BlockTodo
}

/*
	Init functions to create new default todo structs
*/

// Creates a new instance of todo
func newSimpleTodo(title string, task string) (*SimpleTodo, int) {
	if title == "" {
		return nil, NotValidParameters
	}

	return &SimpleTodo{title, task, 0}, NoErr
}

// Creates new instance of ComplexMetadata
func newComplexMetadata(title string, task string) (*ComplexMetadata, int) {
	if title == "" {
		return nil, NotValidCommand
	}

	return &ComplexMetadata{title, task, COMPLEX_VERSION, 0, 0}, NoErr
}

// Creates new instance of BlockTodo
func newBlockTodo() (*BlockTodo, int) {
	return &BlockTodo{[MAX_BLOCK_SIZE]SimpleTodo{}, 0}, NoErr
}

func newComplexTodo(title string, task string) (*ComplexTodo, int) {
	if title == "" {
		return nil, NotValidParameters
	}

	metadata, mErr := newComplexMetadata(title, task)
	if mErr != NoErr {
		return nil, mErr
	}

	block, bErr := newBlockTodo()
	if bErr != NoErr {
		return nil, bErr
	}

	return &ComplexTodo{*metadata, [MAX_COMPLEX_SIZE]BlockTodo{*block}}, NoErr
}
