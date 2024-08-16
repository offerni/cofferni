package models

import "time"

type Item struct {
	Available  bool
	CreatedAt  time.Time
	ID         string `gorm:"primaryKey"`
	ModifiedAt time.Time
	Name       string
}

func (Item) TableName() string {
	return "items"
}
