package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type UserCharacters struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
	UserId string `gorm:"not null" json:"user_id"`
	CharacterMasterId string `gorm:"not null" json:"character_master_id"`
	Level int `gorm:"not null;default:1," json:"level"`
	ExperiencePoint int `gorm:"not null" json:"experience_point"`
	Name string `gorm:"not null" json:"name"`
}

func NewUserCharacters(
    user_id string,
    character_master_id string,
    level int,
    experience_point int,
    name string,
) *UserCharacters {
    return &UserCharacters{
        Id: uuid.NewString(),
        UserId: user_id,
        CharacterMasterId: character_master_id,
        Level: level,
        ExperiencePoint: experience_point,
        Name: name,
    }
}
