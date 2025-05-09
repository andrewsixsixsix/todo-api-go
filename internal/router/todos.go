package router

import (
	"todo-api/internal/handler"
	"todo-api/internal/service"

	"github.com/go-chi/chi/v5"
)

func todoRouter() *chi.Mux {
	r := chi.NewRouter()

	todoService := service.GetTodoService()
	handler := handler.NewTodoHandler(todoService)

	r.Get("/", handler.GetTodos)
	r.Post("/", handler.CreateTodo)
	r.Put("/", handler.UpdateTodo)
	r.Delete("/{id}", handler.DeleteTodo)

	return r
}
