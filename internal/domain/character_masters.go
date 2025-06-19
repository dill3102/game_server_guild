package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CharacterMasters struct {
	gorm.Model
	Id              string `gorm:"column:id;primaryKey" json:"id"`
	Name            string `gorm:"not null" json:"name"`
	Level           int    `gorm:"not null;default:0," json:"level"`
	Reality         int    `gorm:"not null;default:0," json:"reality"`
	ExperiencePoint int    `gorm:"not null;default:0," json:"experience_point"`
	Attack          int    `gorm:"not null;default:0," json:"attack"`
	Defense         int    `gorm:"not null;default:0," json:"defense"`
	Speed           int    `gorm:"not null;default:0," json:"speed"`
}

func NewCharacterMasters(
	name string,
	level int,
	reality int,
	experience_point int,
	attack int,
	defense int,
	speed int,
) *CharacterMasters {
	return &CharacterMasters{
		Id:              uuid.NewString(),
		Name:            name,
		Level:           level,
		Reality:         reality,
		ExperiencePoint: experience_point,
		Attack:          attack,
		Defense:         defense,
		Speed:           speed,
	}
}
