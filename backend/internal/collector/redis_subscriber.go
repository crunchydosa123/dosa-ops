package collector

import (
	"context"
	"encoding/json"
	"log"

	"github.com/crunchydosa123/dosa-ops/internal/models"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func StartCollector() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pubsub := rdb.Subscribe(ctx, "metrics")

	ch := pubsub.Channel()

	for msg := range ch {
		var metric models.Metric

		err := json.Unmarshal([]byte(msg.Payload), &metric)
		if err != nil {
			log.Println("invalid metric: ", err)
			continue
		}

		log.Printf("received metric: %+v\n", metric)
	}
}
