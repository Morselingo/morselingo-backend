package model

import "time"

type Message struct {
	MessageId string    `json:"id"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}

func ValidateMessage(message Message) error {
	return validate.Struct(message)
}
