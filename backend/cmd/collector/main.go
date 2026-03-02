package main

import (
	"log"
	"time"

	"github.com/crunchydosa123/dosa-ops/db"
	"github.com/crunchydosa123/dosa-ops/internal/collector"
)

func main() {

	database := db.Connect()

	for {
		log.Println("running collector...")

		err := collector.Run(database)
		if err != nil {
			log.Println(err)
		}

		time.Sleep(10 * time.Second)
	}
}
