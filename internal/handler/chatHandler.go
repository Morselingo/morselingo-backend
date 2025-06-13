package handler

import (
	"log"
	"net/http"

	"github.com/Morselingo/morselingo-backend/internal/service"
	"github.com/gorilla/websocket"
)

var webSocketUpgrader = websocket.Upgrader{
	//TODO: Remove this (allow cross origin for local development)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHandler struct {
	service service.ChatService
}

func NewChatHandler(service service.ChatService) *ChatHandler {
	return &ChatHandler{service: service}
}

func (handler ChatHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading WebSocket")
		return
	}
	defer conn.Close()
	//TODO: send new messages to the connected clients
}

func (handler ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
	w.WriteHeader(http.StatusAccepted)
}
