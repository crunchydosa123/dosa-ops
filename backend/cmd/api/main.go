package main

import (
	"log"

	"github.com/crunchydosa123/dosa-ops/db"
	"github.com/crunchydosa123/dosa-ops/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {

	database := db.Connect()

	r := gin.Default()

	r.POST("/login", api.Login(database))
	r.POST("/signup", api.Signup(database))

	//apigroup := r.Group("/api")
	//apigroup.Use(api.AuthMiddleware())

	api.RegisterRoutes(r, database)

	log.Println("API running on :8080")

	r.Run(":8080")
}
