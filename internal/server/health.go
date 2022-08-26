package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthRoute commonservice

func (t *HealthRoute) registry() {
	t.router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "health")
	})
}
