package api

import (
	"net/http"

	"github.com/crunchydosa123/dosa-ops/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetMetric(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		service := c.Query("service")
		metric := c.Query("metric")

		data, err := repository.GetMetrics(db, service, metric)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
