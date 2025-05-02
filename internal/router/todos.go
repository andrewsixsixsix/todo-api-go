package router

import (
	"todo-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func TodoRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handler.GetTodos)
	r.Post("/", handler.CreateTodo)
	r.Put("/", handler.UpdateTodo)
	r.Delete("/{id}", handler.DeleteTodo)

	return r
}
