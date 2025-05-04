package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	openapiRouter := openapiRouter()
	todoRouter := todoRouter()

	r.Mount("/api/openapi", openapiRouter)
	r.Mount("/api/todos", todoRouter)

	return r
}
