package collector

import (
	"database/sql"
	"log"

	"github.com/crunchydosa123/dosa-ops/internal/repository"
)

func Run(db *sql.DB) error {

	services, err := repository.GetServices(db)
	if err != nil {
		return err
	}

	for _, service := range services {

		body, err := scrape(service.URL)
		if err != nil {
			log.Println("scrape failed:", err)
			continue
		}

		metrics := parseMetrics(string(body))

		for name, value := range metrics {

			repository.InsertMetric(db, service.Name, name, value)
		}
	}

	return nil
}
