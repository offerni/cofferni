package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/menu"
	"github.com/offerni/cofferni/rest"
	"github.com/offerni/cofferni/sqlite"
	"github.com/offerni/cofferni/sqlite/connection"
	"github.com/offerni/cofferni/sqlite/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		if os.IsNotExist(err) {
			log.Println(".env file not found, using embedded environment variables")
		} else {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	db := initializeDB()
	defer closeDbConnection(db)

	deps := initDependencies(initDependenciesOpts{
		db: db,
	})

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allowing wildcard for now because I'm lazy and this will run only on LAN
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router := chi.NewRouter()
	router.Use(corsMiddleware.Handler)

	server, err := rest.NewServer(rest.NewServerOpts{
		MenuService: deps.menuService,
		Router:      router,
		Port:        port,
	})
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server stopped gracefully")
}

// initializeDB creates the tables in the database if they don't exist already.
func initializeDB() *connection.DB {
	log.Println("Running Migration...")

	db, err := connection.Open(connection.DbConfig{
		Name: fmt.Sprintf("%s.db", os.Getenv("DATABASE_NAME")),
	})
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

type initDependenciesOpts struct {
	db *connection.DB
}

type dependencies struct {
	menuService *menu.Service
}

func initDependencies(deps initDependenciesOpts) *dependencies {
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

	return &dependencies{
		menuService: menuSvc,
	}
}

// seedData populates tables with pre-defined data
func seedData(itemRepo cofferni.ItemRepository) error {
	_, err := itemRepo.CreateAll(context.Background(), cofferni.ItemCreateAllOpts{
		Items: []*cofferni.ItemCreateOpts{
			{Name: "Espresso", Available: true},
			{Name: "Iced Espresso", Available: true},
			{Name: "Americano", Available: true},
			{Name: "Iced Americano", Available: true},
			{Name: "Latte", Available: true},
			{Name: "Flat White", Available: true},
			{Name: "Mocha Latte", Available: true},
			{Name: "Iced Mocha Latte", Available: true},
			{Name: "Cappuccino", Available: true},
			{Name: "Hot Chocolate", Available: true},
			{Name: "Pumpkin Spice Latte", Available: true},
			{Name: "Affogato", Available: true},
			{Name: "(Cocktail) Aperol Spritz", Available: true},
			{Name: "(Cocktail) Espresso Martini", Available: true},
		},
	})

	if err != nil {
		return fmt.Errorf("err seeding items: %v, skipping", err)
	}

	return nil
}
