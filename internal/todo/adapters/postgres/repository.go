package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hamedslyn/heli-todo/internal/todo/domain"
	"github.com/hamedslyn/heli-todo/internal/todo/ports"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) ports.TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, todo domain.TodoItem) (domain.TodoItem, error) {
	query := `INSERT INTO todo_items (description, due_date) VALUES ($1, $2) RETURNING id;`

	if err := r.db.QueryRowContext(ctx, query, todo.Description, todo.DueDate).Scan(&todo.ID); err != nil {
		return domain.TodoItem{}, fmt.Errorf("insert todo: %w", err)
	}

	return todo, nil
}
