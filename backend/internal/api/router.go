package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func RegisterRoutes(r *gin.Engine, db *pgx.Conn) {
	r.GET("/services", GetServices(db))
	r.GET("metrics", GetMetric(db))
}
