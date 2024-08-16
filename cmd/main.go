package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/offerni/cofferni/sqlite/connection"
	"github.com/offerni/cofferni/sqlite/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	migrateDbTables()
}

// migrateDbTables creates the tables in the database if they don't exist already.
func migrateDbTables() {
	log.Println("Running Migration...")

	db, err := connection.Open(connection.DbConfig{Name: fmt.Sprintf("%s.db", os.Getenv("DATABASE_NAME"))})
	if err != nil {
		panic(err)
	}

	log.Println("Creating `items` Table")
	err = db.DB.AutoMigrate(&models.Item{})
	if err != nil {
		panic(err)
	}

	log.Println("Creating `orders` Table")
	err = db.DB.AutoMigrate(&models.Order{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB.DB()
	if err != nil {
		panic(err)
	}

	defer sqlDB.Close()

	log.Print("Migration Completed.")
}
