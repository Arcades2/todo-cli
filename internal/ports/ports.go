package ports

import "todo-cli/internal/core/domain"

type TodoRepository interface {
	GetAll() ([]domain.Todo, error)
	GetByID(id string) (domain.Todo, error)
	Create(todo domain.Todo) (string, error)
	Update(todo domain.Todo) error
	Delete(id string) error
}

type TodoService interface {
	GetAll() ([]domain.Todo, error)
	GetByID(id string) (domain.Todo, error)
	Create(description string) (string, error)
	ToggleDone(id string) error
	Delete(id string) error
}
