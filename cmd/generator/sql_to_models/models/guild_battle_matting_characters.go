package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type GuildBattleMattingCharacters struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
	INDEX time.Time `gorm:"" json:"INDEX"`
}

func NewGuildBattleMattingCharacters(
    INDEX time.Time,
) *GuildBattleMattingCharacters {
    return &GuildBattleMattingCharacters{
        Id: uuid.NewString(),
        INDEX: INDEX,
    }
}
