package model

import "time"

const (
	UserMessageType = "user_message"
	UserJoinedType  = "user_joined"
	UserLeftType    = "user_left"
)

type Message struct {
	Type      string    `json:"type"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

type UserMessage struct {
	Content string `json:"content" validate:"required"`
}

func ValidateUserMessage(message UserMessage) error {
	return validate.Struct(message)
}
