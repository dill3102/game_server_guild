package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type EventHandler interface {
}

type eventHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewEventHandler(
	usersRepository repository.UsersRepository,
) EventHandler {
	return &eventHandlerImpl{
		usersRepository: usersRepository,
	}
}
