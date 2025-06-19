package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type AuthHandler interface {
}

type authHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewAuthHandler(
	usersRepository repository.UsersRepository,
) AuthHandler {
	return &authHandlerImpl{
		usersRepository: usersRepository,
	}
}
