package router

import (
	"todo-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func OpenAPIRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handler.OpenAPI)
	r.Get("/ui", handler.UI)

	return r
}
