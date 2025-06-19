package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type UserHandler interface {
}

type userHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewUserHandler(
	usersRepository repository.UsersRepository,
) UserHandler {
	return &userHandlerImpl{
		usersRepository: usersRepository,
	}
}
