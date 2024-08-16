package models

import "time"

type Order struct {
	ID          string `gorm:"primaryKey"`
	ItemID      string
	Observation *string
	Quantity    uint
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

func (Order) TableName() string {
	return "orders"
}
