package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/crunchydosa123/dosa-ops/internal/models"
	"github.com/redis/go-redis/v9"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	metric := models.Metric{
		Service:   "auth-service",
		Metric:    "cpu",
		Value:     0.65,
		Timestamp: time.Now().Unix(),
	}

	data, _ := json.Marshal(metric)

	rdb.Publish(context.Background(), "metrics", data)
}
