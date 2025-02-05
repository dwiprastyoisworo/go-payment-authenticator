package database

import (
	"context"
	"database/sql"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/models"
)

type AccessToken struct {
}

func NewAccessToken() AccessTokenRepository {
	return &AccessToken{}
}

type AccessTokenRepository interface {
	InsertAccessToken(ctx context.Context, accessToken models.AccessToken, db *sql.Tx) error
}

func (a AccessToken) InsertAccessToken(ctx context.Context, accessToken models.AccessToken, db *sql.Tx) error {
	// query to insert access token
	query := "INSERT INTO access_tokens (token, client_id, expires_at) VALUES ($1, $2, $3)"

	// execute the query
	_, err := db.ExecContext(ctx, query, accessToken.Token, accessToken.ClientID, accessToken.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}
