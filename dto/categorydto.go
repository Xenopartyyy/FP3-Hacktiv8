package dto

import (
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
