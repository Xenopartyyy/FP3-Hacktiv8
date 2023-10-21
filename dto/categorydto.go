package dto

import (
	"FP3-Hacktiv8/model"
	_ "FP3-Hacktiv8/model"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Tasks     []model.Task
}
