package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"product_service/internal/delivery"
	"product_service/internal/repository"
	"product_service/internal/usecase"
)

func connectWithRetry(dbURL string) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("postgres", dbURL)
		if err == nil {
			return db, nil
		}
		log.Printf("Waiting for database... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}
	return nil, err
}

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/product_db?sslmode=disable"
	}

	db, err := connectWithRetry(dbURL)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	log.Println("Connected to DB")

	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)

	router := gin.Default()
	delivery.NewProductHandler(router, productUC)

	log.Println("Product Service running on :8080")
	router.Run(":8080")
}
