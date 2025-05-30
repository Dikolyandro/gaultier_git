package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/da-er-gaultier/order_service/internal/delivery/http"
	"github.com/da-er-gaultier/order_service/internal/repository/postgres"
	"github.com/da-er-gaultier/order_service/internal/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}
	defer db.Close()

	repo := postgres.NewOrderRepo(db)
	uc := usecase.NewOrderUsecase(repo)

	r := gin.Default()
	http.NewOrderHandler(r, uc)

	log.Println("Order Service running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
