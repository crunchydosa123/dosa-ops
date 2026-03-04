package main

import (
	"github.com/crunchydosa123/dosa-ops/internal/scraper"
	"github.com/redis/go-redis/v9"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	scraper.StartScraper(rdb)
}
