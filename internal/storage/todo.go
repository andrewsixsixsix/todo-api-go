package storage

import "database/sql"

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
