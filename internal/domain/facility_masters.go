package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FacilityMasters struct {
	gorm.Model
	Id               string `gorm:"column:id;primaryKey" json:"id"`
	SizeX            int    `gorm:"not null;default:1," json:"size_x"`
	SizeY            int    `gorm:"not null;default:1," json:"size_y"`
	Name             string `gorm:"not null" json:"name"`
	Description      string `gorm:"not null" json:"description"`
	FacilityType     int    `gorm:"not null;default:0," json:"facility_type"`
	ParentFacilityId string `gorm:"" json:"parent_facility_id"`
}

func NewFacilityMasters(
	size_x int,
	size_y int,
	name string,
	description string,
	facility_type int,
	parent_facility_id string,
) *FacilityMasters {
	return &FacilityMasters{
		Id:               uuid.NewString(),
		SizeX:            size_x,
		SizeY:            size_y,
		Name:             name,
		Description:      description,
		FacilityType:     facility_type,
		ParentFacilityId: parent_facility_id,
	}
}
