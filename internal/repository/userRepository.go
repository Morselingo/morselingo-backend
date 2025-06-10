package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(context context.Context, name string, passwordHash string) error
	UserExistsByName(context context.Context, name string) (bool, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (repository userRepository) CreateUser(ctx context.Context, name string, passwordHash string) error {
	panic("unimplemented")
}

func (repository userRepository) UserExistsByName(ctx context.Context, name string) (bool, error) {
	panic("unimplemented")
}
