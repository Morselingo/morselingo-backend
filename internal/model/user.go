package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type RegisterUserInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type User struct {
	Id           int       `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"required,min=3"`
	PasswordHash string    `json:"password" validate:"required,min=8,max=100"`
	CreationTime time.Time `json:"created_at" validate:"required"`
}

var validate = validator.New()

func ValidateRegisterUserInput(userInput RegisterUserInput) error {
	return validate.Struct(userInput)
}

func ValidateUser(user User) error {
	return validate.Struct(user)
}
