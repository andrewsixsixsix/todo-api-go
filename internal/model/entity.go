package model

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
}
