package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatRepository interface {
}

type chatRepository struct {
	db *pgxpool.Pool
}

func NewChatRepository(db *pgxpool.Pool) ChatRepository {
	return &chatRepository{db: db}
}
