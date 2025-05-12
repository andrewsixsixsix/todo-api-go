package router

import (
	mw "todo-api/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(mw.LoggerMiddleware)
	r.Use(middleware.Recoverer)

	openapiRouter := openapiRouter()
	todoRouter := todoRouter()

	r.Mount("/api/openapi", openapiRouter)
	r.Mount("/api/todos", todoRouter)

	return r
}
