package models

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claim struct {
	ClientID string `json:"client_id"`
	jwt.RegisteredClaims
}

func (c *Claim) GenerateToken(secretKey string, expired time.Time) (string, error) {
	c.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expired), // Token berlaku 1 jam
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   c.ClientID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secretKey))
}
