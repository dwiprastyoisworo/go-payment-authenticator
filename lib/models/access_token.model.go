package models

import "time"

type AccessToken struct {
	Token        string    `json:"token"`
	ClientID     uint      `json:"client_id"`
	UserID       string    `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
	Revoked      bool      `json:"revoked"`
	CreatedAt    time.Time `json:"created_at"`
}
