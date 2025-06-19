package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type UserFacilities struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
	UserId string `gorm:"not null" json:"user_id"`
	FacilityMasterId string `gorm:"not null" json:"facility_master_id"`
	Level int `gorm:"not null;default:1," json:"level"`
	Name string `gorm:"not null" json:"name"`
	IsSetting int `gorm:"default:0," json:"is_setting"`
	LastApproveAt time.Time `gorm:"" json:"last_approve_at"`
}

func NewUserFacilities(
    user_id string,
    facility_master_id string,
    level int,
    name string,
    is_setting int,
    last_approve_at time.Time,
) *UserFacilities {
    return &UserFacilities{
        Id: uuid.NewString(),
        UserId: user_id,
        FacilityMasterId: facility_master_id,
        Level: level,
        Name: name,
        IsSetting: is_setting,
        LastApproveAt: last_approve_at,
    }
}
