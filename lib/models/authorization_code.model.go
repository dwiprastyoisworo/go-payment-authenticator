package models

import "time"

type AuthorizationCode struct {
	Code     string `json:"code"`
	ClientID uint   `json:"client_id"`
	//UserID      string    `json:"user_id"`
	ExpiresAt   time.Time `json:"expires_at"`
	Used        bool      `json:"used"`
	RedirectURI string    `json:"redirect_uri"`
	CreatedAt   time.Time `json:"created_at"`
}

type AuthorizationRequest struct {
	ClientID    string `json:"client_id"`
	RedirectURI string `json:"redirect_uri"`
}
