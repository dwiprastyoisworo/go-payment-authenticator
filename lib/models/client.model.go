package models

import "time"

type Client struct {
	ID           uint      `json:"id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	Name         string    `json:"name"`
	RedirectURIs []string  `json:"redirect_uris"`
	Enabled      bool      `json:"enabled"`
	CreatedAt    time.Time `json:"created_at"`
}
