package model

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    int    `json:"priority" validate:"required,min=1,max=3"`
	Status      string `json:"status" validate:"required,len=1"`
}

type UpdateTodoRequest struct {
	ID          TodoID `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    int    `json:"priority" validate:"required,min=1,max=3"`
	Status      string `json:"status" validate:"required,len=1"`
}
