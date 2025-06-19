package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type MeHandler interface {
}

type meHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewMeHandler(
	usersRepository repository.UsersRepository,
) MeHandler {
	return &meHandlerImpl{
		usersRepository: usersRepository,
	}
}
