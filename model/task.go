package model

import (
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(50)" json:"title" validate:"required"`
	Description string `gorm:"type:varchar(100)" json:"description" validate:"required"`
	Status      bool   `json:"status" validate:"boolean"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User     `gorm:"foreignKey:UserID"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}

func (Task) TableName() string {
	return "task"
}
