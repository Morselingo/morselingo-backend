package model

import (
	"time"
)

type User struct {
	Id           int       `json:"id" validate:"required"`
	Username     string    `json:"username" validate:"required,min=3"`
	PasswordHash string    `json:"password" validate:"required,min=8,max=100"`
	CreationTime time.Time `json:"created_at" validate:"required"`
}

func ValidateUser(user User) error {
	return validate.Struct(user)
}
