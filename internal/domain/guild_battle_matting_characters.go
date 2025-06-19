package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuildBattleMattingCharacters struct {
	gorm.Model
	Id                 string `gorm:"column:id;primaryKey" json:"id"`
	MattingId          string `gorm:"not null" json:"matting_id"`
	CharacterId        string `gorm:"not null" json:"character_id"`
	UserId             string `gorm:"not null" json:"user_id"`
	UserCharacterIndex int    `gorm:"not null;default:0," json:"user_character_index"`
	GuildId            string `gorm:"not null" json:"guild_id"`
}

func NewGuildBattleMattingCharacters(
	matting_id string,
	character_id string,
	user_id string,
	user_character_index int,
	guild_id string,
) *GuildBattleMattingCharacters {
	return &GuildBattleMattingCharacters{
		Id:                 uuid.NewString(),
		MattingId:          matting_id,
		CharacterId:        character_id,
		UserId:             user_id,
		UserCharacterIndex: user_character_index,
		GuildId:            guild_id,
	}
}
