package http

import (
	"net/http"

	"github.com/hamedslyn/heli-todo/internal/todo/domain"
	"github.com/hamedslyn/heli-todo/internal/todo/usecase"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	service *usecase.TodoService
}

func NewTodoHandler(service *usecase.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Create(c echo.Context) error {
	var req CreateTodoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
	}

	todo := domain.TodoItem{
		Description: req.Description,
		DueDate:     req.DueDate,
	}

	createdTodo, err := h.service.Create(c.Request().Context(), todo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to create todo: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, createdTodo)
}
