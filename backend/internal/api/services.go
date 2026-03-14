package api

import (
	"net/http"

	"github.com/crunchydosa123/dosa-ops/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetServices(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		services, err := repository.GetServices(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, services)
	}
}
