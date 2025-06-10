package service

import (
	"context"
	"errors"

	"github.com/Morselingo/morselingo-backend/internal/model"
	"github.com/Morselingo/morselingo-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(ctx context.Context, registerUserInput model.RegisterUserInput) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (service *userService) RegisterUser(ctx context.Context, registerUserInput model.RegisterUserInput) error {
	exists, err := service.repository.UserExistsByName(ctx, registerUserInput.Name)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUserInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	err = service.repository.CreateUser(ctx, registerUserInput.Name, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
