package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalessathler/twitterlike/internal/auth"
)

func AuthMiddleware(authService *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Authorization is required"})
			return
		}
		user, err := authService.ValidateToken(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.Set("UserID", user.ID)
		c.Next()
	}
}
