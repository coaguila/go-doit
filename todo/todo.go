package todo

type Todo struct {
	Title string
	Task  string
	Done  bool
}

type TodoSlice []Todo

func createTodo(args []string) (*Todo, int) {
	if len(args) != 2 {
		return nil, 0
	}

	return nil, 0
}
