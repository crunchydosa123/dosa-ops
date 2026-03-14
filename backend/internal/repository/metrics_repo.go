package repository

import (
	"context"
	"time"

	"github.com/crunchydosa123/dosa-ops/internal/models"
	"github.com/jackc/pgx/v5"
)

func InsertMetric(db *pgx.Conn, service string, name string, value float64) error {

	_, err := db.Exec(
		context.Background(),
		`INSERT INTO metrics(service, name, value, timestamp)
		 VALUES($1,$2,$3,$4)`,
		service, name, value, time.Now(),
	)

	return err
}

func GetMetrics(db *pgx.Conn, service string, metric string) ([]models.MetricPoint, error) {
	rows, err := db.Query(
		context.Background(),
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
