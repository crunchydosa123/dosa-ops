package repository

import (
	"database/sql"

	"github.com/crunchydosa123/dosa-ops/internal/models"
)

func GetServices(db *sql.DB) ([]models.Service, error) {

	rows, err := db.Query("SELECT id, name, url FROM services")
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
