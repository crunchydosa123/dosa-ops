package main

import (
	"log"

	"github.com/crunchydosa123/dosa-ops/internal/collector"
)

func main() {
	log.Println("running collector...")
	collector.StartCollector()
}
