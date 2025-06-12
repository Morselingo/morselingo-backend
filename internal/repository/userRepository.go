package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Morselingo/morselingo-backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(context context.Context, name string, passwordHash string) error
	UserExistsByName(context context.Context, name string) (bool, error)
	GetUserByName(context context.Context, name string) (model.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (repository userRepository) CreateUser(ctx context.Context, name string, passwordHash string) error {
	query := `
		INSERT INTO users (username, password_hash, created_at)
		VALUES ($1, $2, $3)
	`
	now := time.Now()

	_, err := repository.db.Exec(ctx, query, name, passwordHash, now)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (repository userRepository) UserExistsByName(ctx context.Context, name string) (bool, error) {
	query := `SELECT id FROM users WHERE username = $1 LIMIT 1`

	var id int64
	err := repository.db.QueryRow(ctx, query, name).Scan(&id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}

	return true, nil
}

func (repository userRepository) GetUserByName(ctx context.Context, name string) (model.User, error) {
	query := `SELECT id, username, hashed_password, created_at FROM users WHERE username = $1`

	var user model.User
	err := repository.db.QueryRow(ctx, query, name).Scan(&user.Id, &user.Name, &user.PasswordHash, &user.CreationTime)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, ErrorUserNotFound
		}
		return model.User{}, fmt.Errorf("failed to get user by name '%s': %w", name, err)
	}

	return user, nil
}
