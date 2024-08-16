package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/offerni/cofferni/menu"
	"github.com/offerni/cofferni/sqlite"
	"github.com/offerni/cofferni/sqlite/connection"
	"github.com/offerni/cofferni/sqlite/models"
	"github.com/offerni/cofferni/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db := initializeDB()

	defer closeDbConnection(db)

	initDependencies(dependencies{
		db: db,
	})
}

// initializeDB creates the tables in the database if they don't exist already.
func initializeDB() *connection.DB {
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

	log.Print("Migration Completed.")

	return db
}

type dependencies struct {
	db *connection.DB
}

func initDependencies(deps dependencies) {
	db := deps.db

	// repos
	itemRepo := sqlite.NewItemRepository(db)
	orderRepo := sqlite.NewOrderRepository(db)

	// services
	menuSvc, err := menu.NewService(menu.NewServiceOpts{
		ItemRepository:  itemRepo,
		OrderRepository: orderRepo,
	})
	if err != nil {
		panic(err)
	}

	// TODO: for testing purposes only
	items, err := menuSvc.ItemList(context.Background())
	if err != nil {
		log.Printf("err fetching items: %v", err)
	}

	spew.Dump("Items", items)

	order, err := menuSvc.PlaceOrder(context.Background(), menu.PlaceOrderOpts{
		ItemID:      "1",
		Observation: utils.Pointer("decaf please"),
		Quantity:    1,
	})
	if err != nil {
		log.Printf("err placing order: %v", err)
	}

	spew.Dump("Order Created!", order)

	orderList, err := menuSvc.OrderList(context.Background())
	if err != nil {
		log.Printf("err fetching order list: %v", err)
	}

	spew.Dump("Order List!", orderList)
}

func closeDbConnection(db *connection.DB) {
	sqlDB, err := db.DB.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("err closing database: %v", err)
	}
}
