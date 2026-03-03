package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/services", GetServices(db))
	r.GET("metrics", GetMetric(db))
}
