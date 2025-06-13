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
			Conn: make(chan model.Message, 25),
		},
		username: username,
	}
}

func (wsc *WebSocketClient) Handle(ctx context.Context) {
	defer func() {
		wsc.service.Unregister(wsc.client)
		wsc.conn.Close()
	}()

	wsc.service.Register(wsc.client)

	go wsc.writePump()
	wsc.readPump()
}

func (wsc *WebSocketClient) readPump() {
	for {
		var message model.Message
		if err := wsc.conn.ReadJSON(&message); err != nil {
			//TODO: handle client disconnect
			log.Println("read error message from client: ", err)
			break
		}
		message.CreatedAt = time.Now()
		message.Username = wsc.username
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
