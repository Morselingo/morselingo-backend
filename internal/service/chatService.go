package service

import (
	"sync"

	"github.com/Morselingo/morselingo-backend/internal/model"
)

type Client struct {
	Conn chan model.Message
}

type ChatService interface {
	Register(client Client)
	Unregister(client Client)
	Broadcast(message model.Message)
}

type chatService struct {
	clients    map[Client]bool
	broadcast  chan model.Message
	register   chan Client
	unregister chan Client
	mu         sync.RWMutex
}

func NewChatService() ChatService {
	service := &chatService{
		clients:    make(map[Client]bool),
		broadcast:  make(chan model.Message),
		register:   make(chan Client),
		unregister: make(chan Client),
	}

	go service.start()
	return service
}

func (service *chatService) start() {
	for {
		select {
		case client := <-service.register:
			service.mu.Lock()
			service.clients[client] = true
			service.mu.Unlock()
		case client := <-service.unregister:
			service.mu.Lock()
			if _, ok := service.clients[client]; ok {
				delete(service.clients, client)
				close(client.Conn)
			}
			service.mu.Unlock()
		case message := <-service.broadcast:
			service.mu.RLock()
			for client := range service.clients {
				select {
				case client.Conn <- message:
					//success
				default:
					service.Unregister(client)
				}
			}
			service.mu.RUnlock()
		}
	}
}

func (service *chatService) Broadcast(message model.Message) {
	service.broadcast <- message
}

func (service *chatService) Register(client Client) {
	service.register <- client
}

func (service *chatService) Unregister(client Client) {
	service.unregister <- client
}
