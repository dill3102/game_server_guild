package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuildMembers struct {
	gorm.Model
	Id      string `gorm:"column:id;primaryKey" json:"id"`
	GuildId int    `gorm:"not null" json:"guild_id"`
	UserId  string `gorm:"not null;unique" json:"user_id"`
}

func NewGuildMembers(
	guild_id int,
	user_id string,
) *GuildMembers {
	return &GuildMembers{
		Id:      uuid.NewString(),
		GuildId: guild_id,
		UserId:  user_id,
	}
}
