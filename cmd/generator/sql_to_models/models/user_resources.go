package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type UserResources struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
}

func NewUserResources(
) *UserResources {
    return &UserResources{
        Id: uuid.NewString(),
    }
}
