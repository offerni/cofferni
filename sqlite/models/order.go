package models

import "time"

type Order struct {
	CreatedAt   time.Time
	ID          string `gorm:"primaryKey"`
	ItemID      string
	ModifiedAt  time.Time
	Observation *string
	Quantity    uint
}

func (Order) TableName() string {
	return "orders"
}
