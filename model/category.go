package model

import (
	"time"
)

type Category struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Type      string `gorm:"type:varchar(200)" json:"type" validate:"required"`
	Tasks     []Task
	CreatedAt time.Time
	UpdatedAt time.Time
}
