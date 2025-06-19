package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type UpgradeCostMasters struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
	ItemId string `gorm:"not null" json:"item_id"`
	Amount int `gorm:"not null;default:1," json:"amount"`
	Name string `gorm:"not null" json:"name"`
	Level int `gorm:"not null;default:0," json:"level"`
	Description string `gorm:"not null" json:"description"`
}

func NewUpgradeCostMasters(
    item_id string,
    amount int,
    name string,
    level int,
    description string,
) *UpgradeCostMasters {
    return &UpgradeCostMasters{
        Id: uuid.NewString(),
        ItemId: item_id,
        Amount: amount,
        Name: name,
        Level: level,
        Description: description,
    }
}
