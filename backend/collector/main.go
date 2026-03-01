package collector

import (
	"io"
	"net/http"
	"time"
)

func StartCollector(services []Service) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		<-ticker.C
		for _, s := range services {
			metrics := scrapeMetrics(s.URL)
			saveMetrics(metrics, s.Name)
		}
	}
}

func scrapeMetrics(url string) map[string]float64 {
	resp, _ := http.Get(url + "/metrics")
	body, _ := io.ReadAll(resp.Body)

	return parseMetrics(body)
}
