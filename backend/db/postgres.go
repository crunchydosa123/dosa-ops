package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func Connect() *pgx.Conn {

	conn, err := pgx.Connect(
		context.Background(),
		"postgresql://neondb_owner:npg_VJXaA6P2oOCe@ep-old-king-adwoxomx-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require",
	)

	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	return conn
}
