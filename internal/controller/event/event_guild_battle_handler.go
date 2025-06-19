package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type EventGuildBattleHandler interface {
}

type eventGuildBattleHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewEventGuildBattleHandler(
	usersRepository repository.UsersRepository,
) EventGuildBattleHandler {
	return &eventGuildBattleHandlerImpl{
		usersRepository: usersRepository,
	}
}
