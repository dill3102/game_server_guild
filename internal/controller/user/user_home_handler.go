package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type UserHomeHandler interface {
}

type userHomeHandlerImpl struct {
	usersRepository          repository.UsersRepository
	userFacilitiesRepository repository.UserFacilitiesRepository
}

func NewUserHomeHandler(
	usersRepository repository.UsersRepository,
	userFacilitiesRepository repository.UserFacilitiesRepository,
) UserHomeHandler {
	return &userHomeHandlerImpl{
		usersRepository:          usersRepository,
		userFacilitiesRepository: userFacilitiesRepository,
	}
}
