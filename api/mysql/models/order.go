package models

import "time"

type Order struct {
	ID           string    `gorm:"type:char(36);not null;primaryKey"`
	ItemID       string    `gorm:"type:char(36);not null;index"`
	CustomerName string    `gorm:"type:varchar(255);not null"`
	Observation  *string   `gorm:"type:text"`
	Quantity     uint      `gorm:"not null;default:1"`
	Fulfilled    bool      `gorm:"not null;default:false"`
	CreatedAt    time.Time `gorm:"type:datetime;autoCreateTime"`
	ModifiedAt   time.Time `gorm:"type:datetime;autoUpdateTime"`

	Item Item `gorm:"foreignKey:ItemID;references:ID"`
}

func (Order) TableName() string {
	return "orders"
}
