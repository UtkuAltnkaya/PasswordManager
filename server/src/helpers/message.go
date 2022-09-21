package helpers

import (
	"github/exmaximum/passwordmanager/src/models"
)

type MessageType interface {
	string | models.User | models.Passwords | *models.User | *models.Passwords | interface{}
}

type Message[T MessageType] struct {
	Data    T    `json:"message"`
	Success bool `json:"success"`
}

func NewMessage[T MessageType](data T, success bool) Message[T] {
	return Message[T]{
		Data:    data,
		Success: success,
	}
}
