package storage

import "database/sql"

type TodoStorage struct {
	storage *sql.DB
}

func NewTodoStorage(storage *sql.DB) *TodoStorage {
	return &TodoStorage{storage: storage}
}
