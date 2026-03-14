package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() *pgx.Conn {

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	conn, err := pgx.Connect(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	return conn
}
