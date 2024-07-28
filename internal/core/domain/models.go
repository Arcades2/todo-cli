package domain

type Todo struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Done        bool   `json:"done"`
}

func (t *Todo) ToggleDone() {
	t.Done = !t.Done
}

type TodoList = map[string]Todo

func NewTodo(id string, description string) Todo {
	return Todo{
		Description: description,
		ID:          id,
		Done:        false,
	}
}
