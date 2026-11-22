package lib

import (
	"errors"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func DatabaseURL() (string, error) {
	_ = godotenv.Load()

	if value := os.Getenv("DATABASE_URL"); value != "" {
		return value, nil
	}
	return "", errors.New("database url not provided in .env")
}

func OpenDB() (*sqlx.DB, error) {
	dbURL, err := DatabaseURL()
	if err != nil {
		return nil, err
	}
	return sqlx.Connect("pgx", dbURL)
}
