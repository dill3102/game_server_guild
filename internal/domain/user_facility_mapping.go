package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFacilityMapping struct {
	gorm.Model
	Id           string    `gorm:"column:id;primaryKey" json:"id"`
	UserId       string    `gorm:"not null" json:"user_id"`
	MappingIndex string    `gorm:"not null" json:"mapping_index"`
	Mapping      time.Time `gorm:"" json:"mapping"`
	IsOver       bool      `gorm:"default:0," json:"is_over"`
}

func NewUserFacilityMapping(
	user_id string,
	mapping_index string,
	mapping time.Time,
	is_over bool,
) *UserFacilityMapping {
	return &UserFacilityMapping{
		Id:           uuid.NewString(),
		UserId:       user_id,
		MappingIndex: mapping_index,
		Mapping:      mapping,
		IsOver:       is_over,
	}
}
