package ports

import (
	"context"

	"github.com/hamedslyn/heli-todo/internal/todo/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo domain.TodoItem) (domain.TodoItem, error)
}
