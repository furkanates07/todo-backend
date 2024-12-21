package models

import (
	"github.com/google/uuid"
)

type TodoStatus string

const (
	PENDING   TodoStatus = "PENDING"
	COMPLETED TodoStatus = "COMPLETED"
)

type Todo struct {
	ID          uuid.UUID   `json:"id"`
	UserID      uuid.UUID   `json:"user_id"`       
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      TodoStatus `json:"status"`         
	CreatedAt   MyTime     `json:"created_at"`
}
