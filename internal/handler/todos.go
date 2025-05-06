package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"todo-api/internal/logger"
	"todo-api/internal/model"

	"github.com/go-chi/chi/v5"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// TODO: todos, err := todoService.findAllTodos()
	todos := []model.Todo{
		model.Todo{
			ID:          1,
			Title:       "ToDo 1",
			Description: "ToDo 1 description",
			DueDate:     "2026-01-01",
			Priority:    3,
			Status:      "T",
		},
		model.Todo{
			ID:          2,
			Title:       "ToDo 2",
			Description: "ToDo 2 description",
			DueDate:     "2026-01-01",
			Priority:    2,
			Status:      "T",
		},
	}
	getTodosRes := model.GetTodosResponse{Todos: todos}
	res, err := json.Marshal(getTodosRes)
	if err != nil {
		logger.Logger.Error("failed to marshal response json", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Logger.Error("failed to read request body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	createTodoReq := model.CreateTodoRequest{}
	if err := json.Unmarshal(body, &createTodoReq); err != nil {
		logger.Logger.Error("failed to unmarshal request json", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: id, err := todoService.createTodo(createTodoReq)
	id := 1
	createTodoRes := model.CreateTodoResponse{ID: id}
	res, err := json.Marshal(&createTodoRes)
	if err != nil {
		logger.Logger.Error("failed to marshal response json", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Logger.Error("failed to read request body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	updateTodoReq := model.UpdateTodoRequest{}
	if err := json.Unmarshal(body, &updateTodoReq); err != nil {
		logger.Logger.Error("failed to unmarshal request body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: err := todoService.updateTodo(updateTodoReq)

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		logger.Logger.Warn("missing id path param", slog.String("url", r.URL.String()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: err := todoService.deleteTodo(id)

	w.WriteHeader(http.StatusNoContent)
}
