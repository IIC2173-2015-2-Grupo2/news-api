package util

import (
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

func expirationTime() time.Duration {
	return time.Minute * 60
}

/*
Token path
*/
func Token(secretPassword string, claims map[string]interface{}) (string, error) {

	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	// Set some claims
	token.Claims = claims
	token.Claims["exp"] = time.Now().Add(expirationTime()).Unix()

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(secretPassword))
}
