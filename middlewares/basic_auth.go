package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok || user != username || pass != password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
