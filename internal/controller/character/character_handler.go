package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type CharacterHandler interface {
}

type characterHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewCharacterHandler(
	usersRepository repository.UsersRepository,
) CharacterHandler {
	return &characterHandlerImpl{
		usersRepository: usersRepository,
	}
}
