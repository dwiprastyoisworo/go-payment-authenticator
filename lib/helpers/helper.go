package helpers

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func HashString(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	return string(bytes), err
}
