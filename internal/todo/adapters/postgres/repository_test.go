package postgres

import (
	"context"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/hamedslyn/heli-todo/internal/todo/domain"
)

func TestTodoRepository_Create_InsertsRowAndReturnsID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repo := NewTodoRepository(db)

	todo := domain.TodoItem{
		Description: "Write tests",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	query := regexp.QuoteMeta("INSERT INTO todo_items (description, due_date) VALUES ($1, $2) RETURNING id;")

	rows := sqlmock.NewRows([]string{"id"}).AddRow("42")
	mock.ExpectQuery(query).
		WithArgs(todo.Description, sqlmock.AnyArg()).
		WillReturnRows(rows)

	created, err := repo.Create(context.Background(), todo)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if created.ID != "42" {
		t.Fatalf("expected id '42', got %q", created.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %v", err)
	}
}
