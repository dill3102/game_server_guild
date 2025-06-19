package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuildBattleMattings struct {
	gorm.Model
	Id            string `gorm:"column:id;primaryKey" json:"id"`
	EventId       string `gorm:"not null" json:"event_id"`
	WinnerGuildId string `gorm:"not null" json:"winner_guild_id"`
	GuildIdA      string `gorm:"not null" json:"guild_id_a"`
	GuildIdB      string `gorm:"not null" json:"guild_id_b"`
}

func NewGuildBattleMattings(
	event_id string,
	winner_guild_id string,
	guild_id_a string,
	guild_id_b string,
) *GuildBattleMattings {
	return &GuildBattleMattings{
		Id:            uuid.NewString(),
		EventId:       event_id,
		WinnerGuildId: winner_guild_id,
		GuildIdA:      guild_id_a,
		GuildIdB:      guild_id_b,
	}
}
