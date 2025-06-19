package controller

import (
	"github.com/dill3102/game_server_guild/internal/repository"
)

type GuildFundHandler interface {
}

type guildFundHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewGuildFundHandler(
	usersRepository repository.UsersRepository,
) GuildFundHandler {
	return &guildFundHandlerImpl{
		usersRepository: usersRepository,
	}
}
