package connection

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Name     string
	Password string
	Port     string
	Username string
}

func Open(config DbConfig) (*DB, error) {
	// Create the MySQL DSN without the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
	)

	// Open a connection to MySQL without specifying a database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	defer db.Close()

	// Check if the database exists, and create it if it doesn't
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to create database %s: %w", config.Name, err)
	}

	// Include the database in the DSN and open the GORM connection
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database, got error %w", err)
	}

	return &DB{DB: gormDB}, nil
}
