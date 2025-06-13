package repository

import (
	"context"

	"github.com/Morselingo/morselingo-backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatRepository interface {
	SaveMessage(ctx context.Context, msg model.Message) error
	GetMessages(ctx context.Context, limit, offset int) ([]model.Message, error)
}

type chatRepository struct {
	db *pgxpool.Pool
}

func NewChatRepository(db *pgxpool.Pool) ChatRepository {
	return &chatRepository{db: db}
}

func (c *chatRepository) GetMessages(ctx context.Context, limit int, offset int) ([]model.Message, error) {
	panic("unimplemented")
}

func (c *chatRepository) SaveMessage(ctx context.Context, msg model.Message) error {
	panic("unimplemented")
}
