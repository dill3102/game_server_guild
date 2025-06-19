package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Guilds struct {
	gorm.Model
	Id    string `gorm:"column:id;primaryKey" json:"id"`
	Name  string `gorm:"" json:"name"`
	Level int    `gorm:"not null;default:0," json:"level"`
}

func NewGuilds(
	name string,
	level int,
) *Guilds {
	return &Guilds{
		Id:    uuid.NewString(),
		Name:  name,
		Level: level,
	}
}
