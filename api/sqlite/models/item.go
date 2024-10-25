package models

import (
	"time"
)

type Item struct {
	ID          string    `gorm:"not null;primaryKey"`
	Name        string    `gorm:"not null;uniqueIndex"`
	Description *string   `gorm:"type:text"`
	Available   bool      `gorm:"not null;default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	ModifiedAt  time.Time `gorm:"autoUpdateTime"`
}

func (Item) TableName() string {
	return "items"
}
