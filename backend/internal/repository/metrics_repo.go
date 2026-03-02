package repository

import (
	"database/sql"
	"time"
)

func InsertMetric(db *sql.DB, service string, name string, value float64) error {

	_, err := db.Exec(
		`INSERT INTO metrics(service, name, value, timestamp)
		 VALUES($1,$2,$3,$4)`,
		service, name, value, time.Now(),
	)

	return err
}
