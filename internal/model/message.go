package model

import "time"

type Message struct {
	MessageId int32     `json:"message_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

func ValidateMessage(message Message) error {
	return validate.Struct(message)
}
