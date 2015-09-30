package middleware

import (
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

/*
JWTAuth JSON WEB TOKEN Auth middleware
See: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
*/
func JWTAuth(secretPassword string) gin.HandlerFunc {
	return jwt.Auth(secretPassword)
}
