package models

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"update"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var Users []Aluno

func ValidateUser(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
