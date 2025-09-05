package usecase

import (
	"context"

	"github.com/hamedslyn/heli-todo/internal/todo/domain"
	"github.com/hamedslyn/heli-todo/internal/todo/ports"
)

type TodoService struct {
	repo      ports.TodoRepository
	validator ports.TodoValidator
}

func NewTodoService(repo ports.TodoRepository, validator ports.TodoValidator) *TodoService {
	return &TodoService{
		repo:      repo,
		validator: validator,
	}
}

func (s *TodoService) Create(ctx context.Context, todo domain.TodoItem) (domain.TodoItem, error) {
	if err := s.validator.ValidateCreate(todo); err != nil {
		return domain.TodoItem{}, err
	}

	return s.repo.Create(ctx, todo)
}
