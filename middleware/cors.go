package middleware

import (
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
CORSMiddleware allow CORS.
See: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
*/
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AbortOnError:     false,
		AllowedOrigins:   []string{"*"}, // TODO: set GUI url
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "HEAD"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false, // TODO: add auth
		MaxAge:           12 * time.Hour,
	})
}
