package models

import "time"

type Order struct {
	ID          string    `gorm:"not null;primaryKey"`
	ItemID      string    `gorm:"not null;index"`
	Observation *string   `gorm:"type:text"`
	Quantity    uint      `gorm:"not null;default:1"`
	Fulfilled   bool      `gorm:"not null;default:0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	ModifiedAt  time.Time `gorm:"autoUpdateTime"`

	Item Item `gorm:"foreignKey:ItemID;references:ID"`
}

func (Order) TableName() string {
	return "orders"
}
