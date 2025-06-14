package util

import (
	"log"
	"net/http"
	"time"

	"github.com/Morselingo/morselingo-backend/internal/model"
	"github.com/Morselingo/morselingo-backend/internal/service"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

var WebSocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//TODO: Remove this (allow cross origin for local development)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketClient struct {
	conn     *websocket.Conn
	service  service.ChatService
	client   service.Client
	username string
}

func NewWebSocketClient(conn *websocket.Conn, chatService service.ChatService, username string) *WebSocketClient {
	return &WebSocketClient{
		conn:    conn,
		service: chatService,
		client: service.Client{
			Conn: make(chan model.Message, 100),
		},
		username: username,
	}
}

func (wsc *WebSocketClient) Handle(ctx context.Context) {
	defer func() {
		wsc.service.Unregister(wsc.client)
		leaveMsg := model.Message{
			Type:      model.UserLeftType,
			Username:  wsc.username,
			Content:   wsc.username + " disconnected.",
			CreatedAt: time.Now(),
		}
		wsc.service.Broadcast(leaveMsg)
		wsc.conn.Close()
	}()

	wsc.service.Register(wsc.client)
	joinMsg := model.Message{
		Type:      model.UserJoinedType,
		Username:  wsc.username,
		Content:   wsc.username + " connected.",
		CreatedAt: time.Now(),
	}
	wsc.service.Broadcast(joinMsg)

	go wsc.writePump()
	wsc.readPump()
}

func (wsc *WebSocketClient) readPump() {
	for {
		var userMessage model.UserMessage
		if err := wsc.conn.ReadJSON(&userMessage); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Read error (unexpected ws close) for user %s: %v", wsc.username, err)
			}
			break
		}

		if err := model.ValidateUserMessage(userMessage); err != nil {
			log.Printf("User '%s' send a ws message not adheerent to expected structure: %v", wsc.username, err)
			continue
		}

		message := model.Message{
			Type:      "user_message",
			Username:  wsc.username,
			CreatedAt: time.Now(),
			Content:   userMessage.Content,
		}
		wsc.service.Broadcast(message)
	}
}

func (wsc *WebSocketClient) writePump() {
	for msg := range wsc.client.Conn {
		if err := wsc.conn.WriteJSON(msg); err != nil {
			log.Println("error writing message to client: ", err)
			break
		}
	}
}
