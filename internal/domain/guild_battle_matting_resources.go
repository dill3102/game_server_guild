package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuildBattleMattingResources struct {
	gorm.Model
	Id           string `gorm:"column:id;primaryKey" json:"id"`
	MattingId    string `gorm:"not null" json:"matting_id"`
	ResourceType string `gorm:"not null" json:"resource_type"`
	GuildId      string `gorm:"not null" json:"guild_id"`
	Amount       int    `gorm:"not null;default:0," json:"amount"`
}

func NewGuildBattleMattingResources(
	matting_id string,
	resource_type string,
	guild_id string,
	amount int,
) *GuildBattleMattingResources {
	return &GuildBattleMattingResources{
		Id:           uuid.NewString(),
		MattingId:    matting_id,
		ResourceType: resource_type,
		GuildId:      guild_id,
		Amount:       amount,
	}
}
