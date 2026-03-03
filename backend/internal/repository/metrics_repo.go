package repository

import (
	"database/sql"
	"time"

	"github.com/crunchydosa123/dosa-ops/internal/models"
)

func InsertMetric(db *sql.DB, service string, name string, value float64) error {

	_, err := db.Exec(
		`INSERT INTO metrics(service, name, value, timestamp)
		 VALUES($1,$2,$3,$4)`,
		service, name, value, time.Now(),
	)

	return err
}

func GetMetrics(db *sql.DB, service string, metric string) ([]models.MetricPoint, error) {
	rows, err := db.Query(
		`SELECT timestamp, value
		FROM metrics
		WHERE service=$1 AND name=$2
		ORDER BY timestamp`,
		service, metric,
	)

	if err != nil {
		return nil, err
	}

	var results []models.MetricPoint

	for rows.Next() {
		var m models.MetricPoint
		rows.Scan(&m.Timestamp, &m.Value)
		results = append(results, m)
	}

	return results, nil
}
