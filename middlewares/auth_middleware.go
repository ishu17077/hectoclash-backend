package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication token provided."})
			c.Abort()
			return
		}
		claims, err := helpers.ValiateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
