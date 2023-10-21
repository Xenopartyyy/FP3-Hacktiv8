package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Full_name string `gorm:"type:varchar(200)" validate:"required"`
	Email     string `gorm:"type:varchar(100);unique" validate:"required,email"`
	Password  string `gorm:"type:varchar(200)" validate:"required,min=6"`
	Role      string `gorm:"type:varchar(10)" validate:"required"`
	Tasks     []Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}
