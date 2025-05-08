package router

import (
	"todo-api/internal/handler"
	"todo-api/internal/service"
	"todo-api/internal/storage"

	"github.com/go-chi/chi/v5"
)

func todoRouter() *chi.Mux {
	r := chi.NewRouter()

	todoStorage := storage.NewTodoStorage(storage.Storage())
	todoService := service.NewTodoService(todoStorage)
	handler := handler.NewTodoHandler(todoService)

	r.Get("/", handler.GetTodos)
	r.Post("/", handler.CreateTodo)
	r.Put("/", handler.UpdateTodo)
	r.Delete("/{id}", handler.DeleteTodo)

	return r
}
