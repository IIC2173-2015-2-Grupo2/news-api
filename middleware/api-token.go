package middleware

import "github.com/gin-gonic/gin"

/*
APIToken filter valid API keys.
*/
func APIToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: validate tokens
		c.Next()
	}
}
