package main

import (
	"log"

	"github.com/crunchydosa123/dosa-ops/db"
	"github.com/crunchydosa123/dosa-ops/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {

	database := db.Connect()

	router := gin.Default()

	api.RegisterRoutes(router, database)

	log.Println("API running on :8080")

	router.Run(":8080")
}
