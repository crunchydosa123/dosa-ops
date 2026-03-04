package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	ID   int
	Name string
	URL  string
}

type MetricEvent struct {
	ServiceID int     `json:"service_id"`
	Service   string  `json:"service"`
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

var ctx = context.Background()

func StartScraper(rdb *redis.Client) {

	ticker := time.NewTicker(10 * time.Second)

	for {
		<-ticker.C

		services := getServices()

		for _, s := range services {
			scrapeService(s, rdb)
		}
	}
}

func scrapeService(service Service, rdb *redis.Client) {

	url := fmt.Sprintf("%s/metrics", service.URL)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var metrics map[string]float64

	err = json.NewDecoder(resp.Body).Decode(&metrics)
	if err != nil {
		return
	}

	for name, value := range metrics {

		event := MetricEvent{
			ServiceID: service.ID,
			Service:   service.Name,
			Metric:    name,
			Value:     value,
			Timestamp: time.Now().Unix(),
		}

		data, _ := json.Marshal(event)

		rdb.Publish(ctx, "metrics", data)
	}
}

func getServices() []Service {

	return []Service{
		{
			ID:   1,
			Name: "posts-service",
			URL:  "http://localhost:8081",
		},
	}
}
