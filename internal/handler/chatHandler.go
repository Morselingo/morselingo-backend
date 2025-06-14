package handler

import (
	"log"
	"net/http"

	"github.com/Morselingo/morselingo-backend/internal/auth"
	"github.com/Morselingo/morselingo-backend/internal/service"
	"github.com/Morselingo/morselingo-backend/internal/util"
)

type ChatHandler struct {
	service service.ChatService
}

func NewChatHandler(service service.ChatService) *ChatHandler {
	return &ChatHandler{service: service}
}

func (handler ChatHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := util.WebSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading WebSocket")
		return
	}

	username, ok := auth.GetUsernameFromContext(r.Context())
	if !ok {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	client := util.NewWebSocketClient(conn, handler.service, username)
	go client.Handle(r.Context())
}
