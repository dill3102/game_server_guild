package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuildResources struct {
	gorm.Model
	Id           string `gorm:"column:id;primaryKey" json:"id"`
	UserId       string `gorm:"not null;unique" json:"user_id"`
	Balance      int    `gorm:"not null" json:"balance"`
	ResourceType int    `gorm:"not null" json:"resource_type"`
}

func NewGuildResources(
	user_id string,
	balance int,
	resource_type int,
) *GuildResources {
	return &GuildResources{
		Id:           uuid.NewString(),
		UserId:       user_id,
		Balance:      balance,
		ResourceType: resource_type,
	}
}
