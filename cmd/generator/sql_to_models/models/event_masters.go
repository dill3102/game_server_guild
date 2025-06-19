package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
)

type EventMasters struct {
	gorm.Model
	Id string `gorm:"column:id;primaryKey" json:"id"`
	Since time.Time `gorm:"" json:"since"`
	Until time.Time `gorm:"" json:"until"`
	Title string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	EventType int `gorm:"not null;default:0," json:"event_type"`
	ParentEventId string `gorm:"" json:"parent_event_id"`
}

func NewEventMasters(
    since time.Time,
    until time.Time,
    title string,
    description string,
    event_type int,
    parent_event_id string,
) *EventMasters {
    return &EventMasters{
        Id: uuid.NewString(),
        Since: since,
        Until: until,
        Title: title,
        Description: description,
        EventType: event_type,
        ParentEventId: parent_event_id,
    }
}
