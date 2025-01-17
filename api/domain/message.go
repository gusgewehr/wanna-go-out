package domain

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Text string `json:"text"`
}

type MessageRepository interface {
	Save(message *Message) (*Message, error)
	Get() (messages []Message)
}

type MessageUseCase interface {
	SaveMessage(message *Message) (*Message, error)
	GetAllMessage() (messages []Message)
}
