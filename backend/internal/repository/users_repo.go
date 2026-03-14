package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetHashedPasswordByEmail(db *pgx.Conn, email string) (string, error) {

	var password string

	err := db.QueryRow(
		context.Background(),
		`SELECT password FROM users WHERE email = $1`,
		email,
	).Scan(&password)

	if err != nil {
		return "", err
	}

	return password, nil
}
