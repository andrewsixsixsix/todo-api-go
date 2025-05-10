package model

type TodoID int64

type Todo struct {
	ID          TodoID `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
}
