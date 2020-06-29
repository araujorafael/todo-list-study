package models

import (
	"time"
)

// Task payload definition
type Task struct {
	ID        uint32     `json:"id"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
