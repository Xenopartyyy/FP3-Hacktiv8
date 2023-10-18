package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Full_name  string `gorm:"type:varchar(200)" validate:"required"`
	Email      string `gorm:"type:varchar(100);unique" validate:"required,email"`
	Password   string `gorm:"type:varchar(200)" validate:"required,min=6"`
	Role       string `gorm:"type:varchar(10)" validate:"required"`
	Tasks      []Task
	Created_at time.Time
	Updated_at time.Time
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
