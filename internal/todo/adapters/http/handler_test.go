package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hamedslyn/heli-todo/internal/todo/adapters/validator"
	"github.com/hamedslyn/heli-todo/internal/todo/domain"
	"github.com/hamedslyn/heli-todo/internal/todo/usecase"
	"github.com/labstack/echo/v4"
)

type inMemoryRepo struct{ items []domain.TodoItem }

func (m *inMemoryRepo) Create(ctx context.Context, todo domain.TodoItem) (domain.TodoItem, error) {
	todo.ID = "mock-id"
	m.items = append(m.items, todo)
	return todo, nil
}

func setupHandler() (*TodoHandler, *inMemoryRepo, *echo.Echo) {
	repo := &inMemoryRepo{}
	v := validator.NewTodoValidator()
	svc := usecase.NewTodoService(repo, v)
	h := NewTodoHandler(svc)
	e := echo.New()
	return h, repo, e
}

func TestCreate_Handler_Success(t *testing.T) {
	h, _, e := setupHandler()

	reqBody := CreateTodoRequest{Description: "Write tests", DueDate: time.Now().Add(24 * time.Hour)}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Create(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}

	var got domain.TodoItem
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}
	if got.ID == "" || got.Description != reqBody.Description {
		t.Fatalf("unexpected response: %+v", got)
	}
}

func TestCreate_Handler_ValidationError(t *testing.T) {
	h, _, e := setupHandler()

	reqBody := CreateTodoRequest{Description: "", DueDate: time.Now().Add(24 * time.Hour)}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = h.Create(c)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500 on validation error mapping, got %d", rec.Code)
	}
}
