package ports

import "github.com/hamedslyn/heli-todo/internal/todo/domain"

type TodoValidator interface {
	ValidateCreate(todo domain.TodoItem) error
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v ValidationError) Error() string {
	return v.Message
}
