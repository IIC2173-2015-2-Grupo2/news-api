package middleware

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

/*
GZIP compress responses
*/
func GZIP() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}
