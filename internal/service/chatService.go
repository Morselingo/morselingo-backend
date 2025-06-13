package service

import (
	"github.com/Morselingo/morselingo-backend/internal/repository"
)

type ChatService interface {
	Register()
	Unregister()
	Broadcast()
}

type chatService struct {
	repository repository.ChatRepository
}

func NewChatService(repository repository.ChatRepository) ChatService {
	return &chatService{repository: repository}
}

// Broadcast implements ChatService.
func (c *chatService) Broadcast() {
	panic("unimplemented")
}

// Register implements ChatService.
func (c *chatService) Register() {
	panic("unimplemented")
}

// Unregister implements ChatService.
func (c *chatService) Unregister() {
	panic("unimplemented")
}
