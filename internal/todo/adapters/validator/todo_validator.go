package validator

import (
	"time"

	"github.com/hamedslyn/heli-todo/internal/todo/domain"
	"github.com/hamedslyn/heli-todo/internal/todo/ports"
)

type TodoValidatorImpl struct{}

func NewTodoValidator() *TodoValidatorImpl {
	return &TodoValidatorImpl{}
}

func (v *TodoValidatorImpl) ValidateCreate(todo domain.TodoItem) error {
	var errors []ports.ValidationError

	if todo.Description == "" {
		errors = append(errors, ports.ValidationError{
			Field:   "description",
			Message: "description cannot be empty",
		})
	}

	now := time.Now()
	if todo.DueDate.Before(now) {
		errors = append(errors, ports.ValidationError{
			Field:   "due_date",
			Message: "due date cannot be in the past",
		})
	}

	if len(errors) > 0 {
		return &ValidationErrors{Errors: errors}
	}

	return nil
}

type ValidationErrors struct {
	Errors []ports.ValidationError
}

func (v *ValidationErrors) Error() string {
	if len(v.Errors) == 0 {
		return "validation failed"
	}
	return v.Errors[0].Error()
}

func (v *ValidationErrors) GetErrors() []ports.ValidationError {
	return v.Errors
}
