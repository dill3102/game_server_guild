package controller

import (
	auth_controller "github.com/dill3102/game_server_guild/internal/controller/auth"
	character_controller "github.com/dill3102/game_server_guild/internal/controller/character"
	event_controller "github.com/dill3102/game_server_guild/internal/controller/event"
	friend_controller "github.com/dill3102/game_server_guild/internal/controller/friend"
	guild_controller "github.com/dill3102/game_server_guild/internal/controller/guild"
	message_controller "github.com/dill3102/game_server_guild/internal/controller/message"
	user_controller "github.com/dill3102/game_server_guild/internal/controller/user"
	repository "github.com/dill3102/game_server_guild/internal/repository"
)

type Handler struct {
	user_controller.MeHandler
	user_controller.UserHandler
	auth_controller.AuthHandler
	character_controller.CharacterHandler
	event_controller.EventGuildBattleHandler
	event_controller.EventGuildMattingHandler
	event_controller.EventHandler
	friend_controller.FriendHandler
	guild_controller.GuildFundHandler
	guild_controller.GuildHandler
	message_controller.MessageHandler
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{
		MeHandler:                user_controller.NewMeHandler(repo.UsersRepository),
		UserHandler:              user_controller.NewUserHandler(repo.UsersRepository),
		AuthHandler:              auth_controller.NewAuthHandler(repo.UsersRepository),
		CharacterHandler:         character_controller.NewCharacterHandler(repo.UsersRepository),
		EventGuildBattleHandler:  event_controller.NewEventGuildBattleHandler(repo.UsersRepository),
		EventGuildMattingHandler: event_controller.NewEventGuildMattingHandler(repo.UsersRepository),
		EventHandler:             event_controller.NewEventHandler(repo.UsersRepository),
		FriendHandler:            friend_controller.NewFriendHandler(repo.UsersRepository),
		GuildFundHandler:         guild_controller.NewGuildFundHandler(repo.UsersRepository),
		GuildHandler:             guild_controller.NewGuildHandler(repo.UsersRepository),
		MessageHandler:           message_controller.NewMessageHandler(repo.UsersRepository),
	}
}
