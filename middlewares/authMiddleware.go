package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Authorization token is required"})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			c.Abort()
			return
		}
		for _, token := range helpers.BlockList {
			if token == clientToken {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Invalid token"})
				c.Abort()
				return
			}
		}
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("uid", claims.UserID)
		c.Next()
	}
}
