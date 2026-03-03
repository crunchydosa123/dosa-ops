package api

import (
	"database/sql"
	"net/http"

	"github.com/crunchydosa123/dosa-ops/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetServices(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		services, err := repository.GetServices(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, services)
	}
}
