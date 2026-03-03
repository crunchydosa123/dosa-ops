package api

import (
	"database/sql"
	"net/http"

	"github.com/crunchydosa123/dosa-ops/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetMetric(db *sql.DB) gin.HandlerFunc {
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
