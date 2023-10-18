package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(100)" json:"title" validate:"required"`
	Description string `gorm:"type:varchar(50)" json:"description" validate:"required"`
	Status      bool   `json:"status" validate:"boolean"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
	Created_at  time.Time
	Updated_at  time.Time
}
