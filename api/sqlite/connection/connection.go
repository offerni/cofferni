package connection

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbConfig struct {
	Name string
}

func Open(config DbConfig) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Name), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &DB{DB: db}, nil
}
