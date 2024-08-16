package models

import (
	"time"
)

type Item struct {
	ID         string    `gorm:"primaryKey"`
	Name       string    `gorm:"not null;uniqueIndex"`
	Available  bool      `gorm:"not null;default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	ModifiedAt time.Time `gorm:"autoUpdateTime"`
}

func (Item) TableName() string {
	return "items"
}
