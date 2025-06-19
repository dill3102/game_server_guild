package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type EventGuildMattingHandler interface {
}

type eventGuildMattingHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewEventGuildMattingHandler(
	usersRepository repository.UsersRepository,
) EventGuildMattingHandler {
	return &eventGuildMattingHandlerImpl{
		usersRepository: usersRepository,
	}
}
