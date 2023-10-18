package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey" json:"id"`
	Type       string `gorm:"type:varchar(200)" json:"type" validate:"required"`
	Tasks      []Task
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
