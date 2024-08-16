package models

import "time"

type Item struct {
	ID         string `gorm:"primaryKey"`
	Name       string
	Available  bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (Item) TableName() string {
	return "items"
}
