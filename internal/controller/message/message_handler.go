package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type MessageHandler interface {
}

type messageHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewMessageHandler(
	usersRepository repository.UsersRepository,
) MessageHandler {
	return &messageHandlerImpl{
		usersRepository: usersRepository,
	}
}
