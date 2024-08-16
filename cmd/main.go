package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/offerni/cofferni"
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

func closeDbConnection(db *connection.DB) {
	sqlDB, err := db.DB.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Panicf("err closing database: %v", err)
	}
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

	err = seedData(itemRepo)
	if err != nil {
		log.Printf("err seeding data: %v", err)
	}

	// TODO: for testing purposes only
	items, err := menuSvc.ItemList(context.Background())
	if err != nil {
		log.Panicf("err fetching items: %v", err)
	}

	spew.Dump("Item List", items)

	order, err := menuSvc.PlaceOrder(context.Background(), menu.PlaceOrderOpts{
		ItemID:      "c074edbd-2984-4136-b48b-a498ab52ab88",
		Observation: utils.Pointer("decaf please"),
		Quantity:    1,
	})
	if err != nil {
		log.Panicf("err placing order: %v", err)
	}

	spew.Dump("Order Placed", order)

	orderList, err := menuSvc.OrderList(context.Background())
	if err != nil {
		log.Panicf("err fetching orders: %v", err)
	}

	spew.Dump("Order List", orderList)
}

// seedData populates tables with pre-defined data
func seedData(itemRepo cofferni.ItemRepository) error {
	_, err := itemRepo.CreateAll(context.Background(), cofferni.ItemCreateAllOpts{
		Items: []*cofferni.ItemCreateOpts{
			{
				Name:      "Espresso",
				Available: true,
			},
			{
				Name:      "Iced Espresso",
				Available: true,
			},
			{
				Name:      "Americano",
				Available: true,
			},
			{
				Name:      "Iced Americano",
				Available: true,
			},
			{
				Name:      "Latte",
				Available: true,
			},
			{
				Name:      "Flat White",
				Available: true,
			},
			{
				Name:      "Mocha Latte",
				Available: true,
			},
			{
				Name:      "Iced Mocha Latte",
				Available: true,
			},
			{
				Name:      "Cappuccino",
				Available: true,
			},
			{
				Name:      "Hot Chocolate",
				Available: true,
			},
		},
	})

	if err != nil {
		return fmt.Errorf("err seeding items: %v, skipping", err)
	}

	return nil
}
