package model

type CreateTodoResponse struct {
	ID int `json:"id"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}
