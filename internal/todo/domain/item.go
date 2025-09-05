package domain

import "time"

type TodoItem struct {
	ID          string
	Description string
	DueDate     time.Time
}
