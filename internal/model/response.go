package model

type CreateTodoResponse struct {
	ID TodoID `json:"id"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type ValidationFail struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

type ValidationFailResponse struct {
	Errors []ValidationFail `json:"errors"`
}
