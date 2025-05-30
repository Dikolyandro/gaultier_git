package main

import (
	"log"
	"os"

	"github.com/da-er-gaultier/admin_service/config"
	httpDelivery "github.com/da-er-gaultier/admin_service/internal/delivery/http"
	"github.com/da-er-gaultier/admin_service/internal/repository"
	"github.com/da-er-gaultier/admin_service/internal/repository/postgres"
	"github.com/da-er-gaultier/admin_service/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := config.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Репозиторий (для продуктов и заказов)
	adminRepo := postgres.New(db)

	// HTTP-клиент для user_service
	userClient := repository.NewUserHTTPClient(os.Getenv("USER_SERVICE_URL"))

	// Юзкейс с двухсторонней логикой
	uc := usecase.New(adminRepo, userClient)

	// HTTP-сервер
	r := gin.Default()
	httpDelivery.New(r, uc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Admin Service running on :%s", port)
	r.Run(":" + port)
}
