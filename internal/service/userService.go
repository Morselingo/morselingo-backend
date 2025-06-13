package service

import (
	"context"
	"errors"

	"github.com/Morselingo/morselingo-backend/internal/auth"
	"github.com/Morselingo/morselingo-backend/internal/model"
	"github.com/Morselingo/morselingo-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(ctx context.Context, registerRequest model.RegisterRequest) error
	LoginUser(ctx context.Context, loginRequest model.LoginRequest) (string, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (service *userService) RegisterUser(ctx context.Context, registerRequest model.RegisterRequest) error {
	exists, err := service.repository.UserExistsByName(ctx, registerRequest.Username)
	if err != nil {
		return err
	}
	if exists {
		return ErrorUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrorFailedToHashPassword
	}

	if err := service.repository.CreateUser(ctx, registerRequest.Username, string(hashedPassword)); err != nil {
		return ErrorCreateUserFailed
	}

	return nil
}

func (service *userService) LoginUser(ctx context.Context, loginRequest model.LoginRequest) (string, error) {
	user, err := service.repository.GetUserByName(ctx, loginRequest.Username)
	if err != nil {
		if errors.Is(err, repository.ErrorUserNotFound) {
			return "", ErrorUserNotFound
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password)); err != nil {
		return "", ErrorAuthenticationFailed
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
