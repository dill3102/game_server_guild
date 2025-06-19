package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type GuildResrouces struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
}

func NewGuildResrouces(
) *GuildResrouces {
    return &GuildResrouces{
        Id: uuid.NewString(),
    }
}
