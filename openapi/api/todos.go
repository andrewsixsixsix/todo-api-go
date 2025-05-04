package api

import (
	"fmt"
	"net/http"
	"todo-api/internal/model"

	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
)

const (
	tag   = "todos"
	route = "/api/todos"
)

func TodosOpenAPI(r *openapi3.Reflector) error {
	getAllTodosOp, err := getAllTodosOperation(r)
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return err
	}

	createTodoOp, err := createTodoOperation(r)
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return err
	}

	updateTodoOp, err := updateTodoOperation(r)
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return err
	}

	deleteTodoOp, err := deleteTodoOperation(r)
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return err
	}

	r.AddOperation(getAllTodosOp)
	r.AddOperation(createTodoOp)
	r.AddOperation(updateTodoOp)
	r.AddOperation(deleteTodoOp)

	return nil
}

func getAllTodosOperation(r *openapi3.Reflector) (openapi.OperationContext, error) {
	op, err := r.NewOperationContext(http.MethodGet, route)
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return nil, err
	}

	op.SetSummary("Get all todos")
	op.SetTags(tag)
	op.AddRespStructure(new(GetTodosResponse))

	return op, nil
}

func createTodoOperation(r *openapi3.Reflector) (openapi.OperationContext, error) {
	op, err := r.NewOperationContext(http.MethodPost, route)
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return nil, err
	}

	op.SetSummary("Create todo")
	op.SetTags(tag)
	op.AddReqStructure(new(CreateTodoRequest))
	op.AddRespStructure(new(CreateTodoResponse), func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusCreated
	})

	return op, nil
}

func updateTodoOperation(r *openapi3.Reflector) (openapi.OperationContext, error) {
	op, err := r.NewOperationContext(http.MethodPut, route)
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return nil, err
	}

	op.SetSummary("Update todo")
	op.SetTags(tag)
	op.AddReqStructure(new(UpdateTodoRequest))

	return op, nil
}

func deleteTodoOperation(r *openapi3.Reflector) (openapi.OperationContext, error) {
	op, err := r.NewOperationContext(http.MethodDelete, route+"/{id}")
	if err != nil {
		fmt.Println("failed to create openapi operation context for todos")
		return nil, err
	}

	op.SetSummary("Delete todo")
	op.SetTags(tag)
	op.AddReqStructure(new(DeleteTodoRequest))

	return op, nil
}

type GetTodosResponse struct {
	Todos []model.Todo `json:"todos"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
}

type CreateTodoResponse struct {
	ID int `json:"id"`
}

type UpdateTodoRequest struct {
	ID          int    `json:"id" required:"true"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
}

type DeleteTodoRequest struct {
	ID int `path:"id" required:"true"`
}
