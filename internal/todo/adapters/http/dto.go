package http

import "time"

type CreateTodoRequest struct {
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type TodoResponse struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
