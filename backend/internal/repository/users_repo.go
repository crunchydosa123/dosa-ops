package repository

import (
	"database/sql"
)

func GetHashedPasswordByEmail(db *sql.DB, email string) (string, error) {

	var password string

	err := db.QueryRow(
		`SELECT password FROM users WHERE id = $1`,
		email,
	).Scan(&password)

	if err != nil {
		return "", err
	}

	return password, nil
}
