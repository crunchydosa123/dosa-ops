package repository

import (
	"context"

	"github.com/crunchydosa123/dosa-ops/internal/models"
	"github.com/jackc/pgx/v5"
)

func GetServices(db *pgx.Conn) ([]models.Service, error) {

	rows, err := db.Query(context.Background(), "SELECT id, name, url FROM services")
	if err != nil {
		return nil, err
	}

	var services []models.Service

	for rows.Next() {
		var s models.Service

		rows.Scan(&s.ID, &s.Name, &s.URL)

		services = append(services, s)
	}

	return services, nil
}
