package middleware

import (
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
CORS allow CORS.
See: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
*/
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AbortOnError:    false,
		AllowAllOrigins: true,
		// AllowedOrigins:   []string{"*"}, // TODO: set GUI url
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "HEAD"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
