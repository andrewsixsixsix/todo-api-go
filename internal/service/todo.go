package service

import (
	"todo-api/internal/model"
	"todo-api/internal/storage"
)

type TodoService struct {
	ts *storage.TodoStorage
}

func NewTodoService(ts *storage.TodoStorage) *TodoService {
	return &TodoService{ts: ts}
}

// TODO: return custom errors

func (ts *TodoService) FindAllTodos() ([]model.Todo, error) {
	todos := []model.Todo{
		model.Todo{
			ID:          1,
			Title:       "ToDo 1",
			Description: "ToDo 1 description",
			DueDate:     "2026-01-01",
			Priority:    3,
			Status:      "T",
		},
		model.Todo{
			ID:          2,
			Title:       "ToDo 2",
			Description: "ToDo 2 description",
			DueDate:     "2026-01-01",
			Priority:    2,
			Status:      "T",
		},
	}

	return todos, nil
}

func (ts *TodoService) CreateTodo(todo model.CreateTodoRequest) (int, error) {
	return 1, nil
}

func (ts *TodoService) UpdateTodo(todo model.UpdateTodoRequest) error {
	return nil
}

func (ts *TodoService) DeleteTodo(id int64) error {
	return nil
}
