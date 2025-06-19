package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type UserFacilityMapping struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
}

func NewUserFacilityMapping(
) *UserFacilityMapping {
    return &UserFacilityMapping{
        Id: uuid.NewString(),
    }
}
