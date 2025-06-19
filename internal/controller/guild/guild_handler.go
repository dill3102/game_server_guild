package controller

import (
	"net/http"

	"github.com/dill3102/game_server_guild/internal/repository"
	"github.com/labstack/echo/v4"
)

type GuildHandler interface {
	ApproveJoinGuildReq(ctx echo.Context, guildId string) error
}

type guildHandlerImpl struct {
	usersRepository repository.UsersRepository
}

func NewGuildHandler(
	usersRepository repository.UsersRepository,
) GuildHandler {
	return &guildHandlerImpl{
		usersRepository: usersRepository,
	}
}

func (g *guildHandlerImpl) ApproveJoinGuildReq(ctx echo.Context, guildId string) error {
	// TODO: 実装前の仮置き
	return ctx.JSON(http.StatusNotImplemented, map[string]string{
		"message": "not implemented yet",
	})
}
