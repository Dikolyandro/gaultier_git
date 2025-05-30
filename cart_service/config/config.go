package config

import (
	"os"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
	return &Config{
		DBUrl: getEnv("DB_URL", "postgres://postgres:postgres@localhost:5435/cart_db?sslmode=disable"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
