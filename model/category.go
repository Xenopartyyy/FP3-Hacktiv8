package model

import (
	"time"

	"github.com/go-playground/validator/v10"
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

func (c *Category) ValidateCategory() error {
	validate := validator.New()
	return validate.Struct(c)
}
