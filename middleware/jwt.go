package middleware

import (
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

/*
AuthMiddleware allow CORS.
See: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
*/
func AuthMiddleware(secretPassword string) gin.HandlerFunc {
	return jwt.Auth(secretPassword)
}

/*
Token path
*/
func Token(secretPassword string, claims map[string]interface{}) (string, error) {

	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	// Set some claims
	token.Claims = claims
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(secretPassword))
}
