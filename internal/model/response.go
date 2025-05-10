package model

type CreateTodoResponse struct {
	ID TodoID `json:"id"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}
