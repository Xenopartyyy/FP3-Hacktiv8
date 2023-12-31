package dto

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Full_name string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdate struct {
	ID        uint      `json:"id"`
	Full_name string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}
