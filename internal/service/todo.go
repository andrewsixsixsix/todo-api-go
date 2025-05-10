package service

import (
	"todo-api/internal/model"
	"todo-api/internal/storage"
)

var todoService *TodoService

type TodoService struct {
	todoStorage *storage.TodoStorage
}

func InitTodoService(todoStorage *storage.TodoStorage) {
	todoService = &TodoService{todoStorage: todoStorage}
}

func GetTodoService() *TodoService {
	return todoService
}

// TODO: return custom errors

func (ts *TodoService) FindAllTodos() ([]model.Todo, error) {
	todos, err := ts.todoStorage.FindAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (ts *TodoService) CreateTodo(ctr model.CreateTodoRequest) (model.TodoID, error) {
	todo := model.Todo{
		Title:       ctr.Title,
		Description: ctr.Description,
		DueDate:     ctr.DueDate,
		Priority:    ctr.Priority,
		Status:      ctr.Status,
	}

	id, err := ts.todoStorage.CreateTodo(todo)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ts *TodoService) UpdateTodo(utr model.UpdateTodoRequest) error {
	todo := model.Todo{
		ID:          utr.ID,
		Title:       utr.Title,
		Description: utr.Description,
		DueDate:     utr.DueDate,
		Priority:    utr.Priority,
		Status:      utr.Status,
	}

	if err := ts.todoStorage.UpdateTodo(todo); err != nil {
		return err
	}

	return nil
}

func (ts *TodoService) DeleteTodo(id model.TodoID) error {
	if err := ts.todoStorage.DeleteTodo(id); err != nil {
		return err
	}

	return nil
}
