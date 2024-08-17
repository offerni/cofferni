package models

import (
	"time"
)

type Item struct {
	ID          string    `gorm:"type:char(36);not null;primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	Description *string   `gorm:"type:text"`
	Available   bool      `gorm:"not null;default:false"`
	CreatedAt   time.Time `gorm:"type:datetime;autoCreateTime"`
	ModifiedAt  time.Time `gorm:"type:datetime;autoUpdateTime"`
}

func (Item) TableName() string {
	return "items"
}
