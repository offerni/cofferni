package connection

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &DB{DB: db}, nil
}
