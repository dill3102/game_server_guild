package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type FriendHandler interface {
}

type friendHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewFriendHandler(
	usersRepository repository.UsersRepository,

) FriendHandler {
	return &friendHandlerImpl{
		usersRepository: usersRepository,
	}
}
