package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("failed to marshal response json")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("failed to read request body")
		return
	}

	defer r.Body.Close()

	createTodoReq := model.CreateTodoRequest{}
	if err := json.Unmarshal(body, &createTodoReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("failed to unmarshal request json")
		return
	}

	// TODO: id, err := todoService.createTodo(createTodoReq)
	id := 1
	createTodoRes := model.CreateTodoResponse{ID: id}
	res, err := json.Marshal(&createTodoRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("failed to marshal response json")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("failed to read request body")
		return
	}

	defer r.Body.Close()

	updateTodoReq := model.UpdateTodoRequest{}
	if err := json.Unmarshal(body, &updateTodoReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("failed to unmarshal request body")
		return
	}

	// TODO: err := todoService.updateTodo(updateTodoReq)

	w.WriteHeader(http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("missing id path param")
		return
	}

	// TODO: err := todoService.deleteTodo(id)

	w.WriteHeader(http.StatusOK)
}
