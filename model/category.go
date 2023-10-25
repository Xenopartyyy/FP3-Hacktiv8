package model

import (
	"time"
)

type Category struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Type      string `gorm:"type:varchar(200)" json:"type" validate:"required"`
	Task      []Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Category) TableName() string {
	return "category"
}
