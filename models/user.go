package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt MyTime   `json:"created_at"`
}
