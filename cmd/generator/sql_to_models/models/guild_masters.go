package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type GuildMasters struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
	Level int `gorm:"not null;default:0," json:"level"`
	MaxMembers int `gorm:"not null;default:0," json:"max_members"`
}

func NewGuildMasters(
    level int,
    max_members int,
) *GuildMasters {
    return &GuildMasters{
        Id: uuid.NewString(),
        Level: level,
        MaxMembers: max_members,
    }
}
