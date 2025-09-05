package server

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/hamedslyn/heli-todo/internal/todo/adapters/http"
	"github.com/hamedslyn/heli-todo/internal/todo/adapters/postgres"
	"github.com/hamedslyn/heli-todo/internal/todo/adapters/validator"
	"github.com/hamedslyn/heli-todo/internal/todo/usecase"
	"github.com/hamedslyn/heli-todo/pkg/config"
)

type Server struct {
	echo *echo.Echo
	cfg  *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	db, err := sql.Open("pgx", cfg.Postgres.ConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	todoRepo := postgres.NewTodoRepository(db)
	todoValidator := validator.NewTodoValidator()
	todoService := usecase.NewTodoService(todoRepo, todoValidator)
	todoHandler := http.NewTodoHandler(todoService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	http.RegisterTodoRoutes(e, todoHandler)

	return &Server{
		echo: e,
		cfg:  cfg,
	}
}

func (s *Server) Run() {
	log.Printf("🚀 Server starting on port %s", s.cfg.Port)
	if err := s.echo.Start(":" + s.cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
