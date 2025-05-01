package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/api/todos", handler)
	r.Post("/api/todos", handler)
	r.Put("/api/todos", handler)
	r.Delete("/api/todos", handler)

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {}
