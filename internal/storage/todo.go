package storage

import (
	"database/sql"
	"todo-api/internal/model"
)

var todoStorage *TodoStorage

type TodoStorage struct {
	storage *sql.DB
}

func InitTodoStorage(storage *sql.DB) {
	todoStorage = &TodoStorage{storage: storage}
}

func GetTodoStorage() *TodoStorage {
	return todoStorage
}

func (ts *TodoStorage) FindAll() ([]model.Todo, error) {
	todos := []model.Todo{}

	rows, err := ts.storage.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.Priority, &todo.Status); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (ts *TodoStorage) CreateTodo(todo model.Todo) (model.TodoID, error) {
	_, err := ts.storage.Exec(
		"INSERT INTO todo (title, description, due_date, priority, status) VALUES ($1, $2, $3, $4, $5)",
		todo.Title, todo.Description, todo.DueDate, todo.Priority, todo.Status,
	)
	if err != nil {
		return 0, err
	}

	var id model.TodoID
	row := ts.storage.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ts *TodoStorage) UpdateTodo(todo model.Todo) error {
	_, err := ts.storage.Exec(
		"UPDATE todo SET title = $1, description = $2, due_date = $3, priority = $4, status = $5 WHERE id = $6",
		todo.Title, todo.Description, todo.DueDate, todo.Priority, todo.Status, todo.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TodoStorage) DeleteTodo(id model.TodoID) error {
	_, err := ts.storage.Exec("DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
