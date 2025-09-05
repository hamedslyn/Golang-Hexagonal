package http

import "github.com/labstack/echo/v4"

func RegisterTodoRoutes(e *echo.Echo, h *TodoHandler) {
	api := e.Group("/api/v1")

	todos := api.Group("/todos")
	todos.POST("", h.Create)

	e.GET("/health", func(c echo.Context) error { return c.NoContent(200) })
}
