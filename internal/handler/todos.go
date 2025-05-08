package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"todo-api/internal/logger"
	"todo-api/internal/model"
	"todo-api/internal/service"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	ts *service.TodoService
}

func NewTodoHandler(ts *service.TodoService) *TodoHandler {
	return &TodoHandler{ts: ts}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.ts.FindAllTodos()
	if err != nil {
		logger.Logger().Error("failed to find all todos", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	getTodosRes := model.GetTodosResponse{Todos: todos}
	res, err := json.Marshal(getTodosRes)
	if err != nil {
		logger.Logger().Error("failed to marshal response json", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Logger().Error("failed to read request body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	createTodoReq := model.CreateTodoRequest{}
	if err := json.Unmarshal(body, &createTodoReq); err != nil {
		logger.Logger().Error("failed to unmarshal request json", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.ts.CreateTodo(createTodoReq)
	if err != nil {
		logger.Logger().Error("failed to create todo", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createTodoRes := model.CreateTodoResponse{ID: id}
	res, err := json.Marshal(&createTodoRes)
	if err != nil {
		logger.Logger().Error("failed to marshal response json", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Logger().Error("failed to read request body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	updateTodoReq := model.UpdateTodoRequest{}
	if err := json.Unmarshal(body, &updateTodoReq); err != nil {
		logger.Logger().Error("failed to unmarshal request body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ts.UpdateTodo(updateTodoReq)
	if err != nil {
		logger.Logger().Error("failed to update todo", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		logger.Logger().Warn("missing id path param", slog.String("url", r.URL.String()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, strconv.IntSize)
	if err != nil {
		logger.Logger().Warn("failed to convert id string to int", slog.String("id", idStr))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ts.DeleteTodo(id)
	if err != nil {
		logger.Logger().Warn("failed to delete todo", slog.String("id", idStr))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
