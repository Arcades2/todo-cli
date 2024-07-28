package todorepo

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"todo-cli/internal/core/domain"
)

type jsonRepository struct{}

func NewJSONRepository() *jsonRepository {
	return &jsonRepository{}
}

func (repo *jsonRepository) GetByID(id string) (domain.Todo, error) {
	file, err := os.Open("db/db.json")
	if err != nil {
		return domain.Todo{}, errors.New("failed to open file")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return domain.Todo{}, errors.New("failed to read db")
	}

	var todos domain.TodoList
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return domain.Todo{}, errors.New("failed to unmarshal db")
	}

	todo, ok := todos[id]

	if !ok {
		return domain.Todo{}, errors.New("todo not found")
	}

	return todo, nil
}

func (repo *jsonRepository) GetAll() ([]domain.Todo, error) {
	file, err := os.Open("db/db.json")
	if err != nil {
		return []domain.Todo{}, errors.New("failed to open file")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return []domain.Todo{}, errors.New("failed to read db")
	}

	var todoList domain.TodoList
	if err := json.Unmarshal(bytes, &todoList); err != nil {
		return []domain.Todo{}, errors.New("failed to unmarshal db")
	}

	todos := make([]domain.Todo, len(todoList))
	for id, todo := range todoList {
		todo.ID = id
		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *jsonRepository) Create(todo domain.Todo) (string, error) {
	file, err := os.Open("db/db.json")
	if err != nil {
		return "", errors.New("failed to open file")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", errors.New("failed to read db")
	}

	var todoList domain.TodoList
	if err := json.Unmarshal(bytes, &todoList); err != nil {
		return "", errors.New("failed to unmarshal db")
	}

	todoList[todo.ID] = todo

	newBytes, err := json.Marshal(todoList)
	if err != nil {
		return "", errors.New("failed to marshal db")
	}

	if err := os.WriteFile("db/db.json", newBytes, 0644); err != nil {
		return "", errors.New("failed to write db")
	}

	return todo.ID, nil
}

func (repo *jsonRepository) Delete(id string) error {
	file, err := os.Open("db/db.json")
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return errors.New("failed to read db")
	}

	var todoList domain.TodoList
	if err := json.Unmarshal(bytes, &todoList); err != nil {
		return errors.New("failed to unmarshal db")
	}

	delete(todoList, id)

	newBytes, err := json.Marshal(todoList)
	if err != nil {
		return errors.New("failed to marshal db")
	}

	if err := os.WriteFile("db/db.json", newBytes, 0644); err != nil {
		return errors.New("failed to write db")
	}

	return nil
}

func (repo *jsonRepository) Update(todo domain.Todo) error {
	file, err := os.Open("db/db.json")
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return errors.New("failed to read db")
	}

	var todoList domain.TodoList
	if err := json.Unmarshal(bytes, &todoList); err != nil {
		return errors.New("failed to unmarshal db")
	}

	todoList[todo.ID] = todo

	newBytes, err := json.Marshal(todoList)
	if err != nil {
		return errors.New("failed to marshal db")
	}

	if err := os.WriteFile("db/db.json", newBytes, 0644); err != nil {
		return errors.New("failed to write db")
	}

	return nil
}
