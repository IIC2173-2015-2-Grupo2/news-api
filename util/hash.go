package util

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 10
)

/*
HashPass apply hashing
*/
func HashPass(pass string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

/*
ValidatePass Compare password and the hashed password.
*/
func ValidatePass(pass, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)); err != nil {
		return err
	}
	return nil
}
