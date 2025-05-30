package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DBUrl string
}

func Load() Config {
	return Config{
		DBUrl: os.Getenv("DB_URL"),
	}
}

func ConnectDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to DB")
	return db, nil
}
