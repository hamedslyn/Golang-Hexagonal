package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/hamedslyn/heli-todo/internal/todo/adapters/validator"
	"github.com/hamedslyn/heli-todo/internal/todo/domain"
)

type mockRepo struct {
	items []domain.TodoItem
}

func (m *mockRepo) Create(ctx context.Context, todo domain.TodoItem) (domain.TodoItem, error) {
	todo.ID = "mock-id"
	m.items = append(m.items, todo)
	return todo, nil
}

func (m *mockRepo) last() (domain.TodoItem, bool) {
	if len(m.items) == 0 {
		return domain.TodoItem{}, false
	}
	return m.items[len(m.items)-1], true
}

func TestCreateTodo_Success(t *testing.T) {
	repo := &mockRepo{}
	todoValidator := validator.NewTodoValidator()
	service := NewTodoService(repo, todoValidator)

	todo := domain.TodoItem{
		Description: "Write tests",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	created, err := service.Create(context.Background(), todo)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if created.ID != "mock-id" {
		t.Errorf("expected ID 'mock-id', got %s", created.ID)
	}
	if created.Description != todo.Description {
		t.Errorf("expected description %s, got %s", todo.Description, created.Description)
	}

	last, ok := repo.last()
	if !ok {
		t.Fatalf("expected item persisted in mock repository")
	}
	if last.Description != todo.Description {
		t.Fatalf("expected persisted description %q, got %q", todo.Description, last.Description)
	}
}

func TestCreateTodo_FailsOnPastDueDate(t *testing.T) {
	repo := &mockRepo{}
	todoValidator := validator.NewTodoValidator()
	service := NewTodoService(repo, todoValidator)

	todo := domain.TodoItem{
		Description: "Write tests",
		DueDate:     time.Now().Add(-24 * time.Hour),
	}

	created, err := service.Create(context.Background(), todo)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}

	ve, ok := err.(*validator.ValidationErrors)
	if !ok {
		t.Fatalf("expected error of type *validator.ValidationErrors, got %T", err)
	}

	foundDueDateError := false
	for _, e := range ve.GetErrors() {
		if e.Field == "due_date" {
			foundDueDateError = true
			break
		}
	}
	if !foundDueDateError {
		t.Fatalf("expected validation error for field 'due_date', got: %+v", ve.GetErrors())
	}

	if created != (domain.TodoItem{}) {
		t.Fatalf("expected zero value todo on error, got: %+v", created)
	}
}

func TestCreateTodo_FailsOnEmptyDescription(t *testing.T) {
	repo := &mockRepo{}
	todoValidator := validator.NewTodoValidator()
	service := NewTodoService(repo, todoValidator)

	todo := domain.TodoItem{
		Description: "",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	created, err := service.Create(context.Background(), todo)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}

	ve, ok := err.(*validator.ValidationErrors)
	if !ok {
		t.Fatalf("expected error of type *validator.ValidationErrors, got %T", err)
	}

	foundDescriptionError := false
	for _, e := range ve.GetErrors() {
		if e.Field == "description" {
			foundDescriptionError = true
			break
		}
	}
	if !foundDescriptionError {
		t.Fatalf("expected validation error for field 'description', got: %+v", ve.GetErrors())
	}

	if created != (domain.TodoItem{}) {
		t.Fatalf("expected zero value todo on error, got: %+v", created)
	}
}
