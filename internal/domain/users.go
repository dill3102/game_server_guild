package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id           string `gorm:"column:id;primaryKey" json:"id"`
	Name         string `gorm:"" json:"name"`
	Level        int    `gorm:"not null" json:"level"`
	Introduction string `gorm:"not null" json:"introduction"`
}

func NewUsers(
	name string,
	level int,
	introduction string,
) *Users {
	return &Users{
		Id:           uuid.NewString(),
		Name:         name,
		Level:        level,
		Introduction: introduction,
	}
}
