package todosrv

import (
	"errors"
	"todo-cli/internal/core/domain"
	"todo-cli/internal/pkg/uidgen"
	"todo-cli/internal/ports"
)

type service struct {
	todoRepository ports.TodoRepository
	uidGen         uidgen.UIDGen
}

func New(todoRepository ports.TodoRepository, uidGen uidgen.UIDGen) *service {
	return &service{
		todoRepository: todoRepository,
		uidGen:         uidGen,
	}
}

func (srv *service) GetAll() ([]domain.Todo, error) {
	todos, err := srv.todoRepository.GetAll()
	if err != nil {
		return []domain.Todo{}, err
	}

	return todos, nil
}

func (srv *service) GetByID(id string) (domain.Todo, error) {
	todo, err := srv.todoRepository.GetByID(id)
	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (srv *service) Create(description string) (string, error) {
	if description == "" {
		return "", errors.New("description cannot be empty")
	}

	todo := domain.NewTodo(srv.uidGen.New(), description)
	id, err := srv.todoRepository.Create(todo)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (srv *service) ToggleDone(id string) error {
	todo, err := srv.todoRepository.GetByID(id)
	if err != nil {
		return errors.New("todo not found")
	}

	todo.ToggleDone()

	if err := srv.todoRepository.Update(todo); err != nil {
		return errors.New("failed to update todo")
	}

	return nil
}

func (srv *service) Delete(id string) error {
	if err := srv.todoRepository.Delete(id); err != nil {
		return errors.New("failed to delete todo")
	}

	return nil
}
