package database

import (
	"context"
	"database/sql"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/models"
)

type AuthorizationCode struct {
}

func NewAuthorizationCodeRepository() AuthorizationCodeRepository {
	return &AuthorizationCode{}
}

type AuthorizationCodeRepository interface {
	GetAuthorizationCode(ctx context.Context, code string, db *sql.Conn) (*models.AuthorizationCode, error)
	InsertAuthorizationCode(ctx context.Context, authCode models.AuthorizationCode, db *sql.Conn) error
}

// GetAuthorizationCode is a function to get authorization code
func (a AuthorizationCode) GetAuthorizationCode(ctx context.Context, code string, db *sql.Conn) (*models.AuthorizationCode, error) {

	// query to get authorization code
	query := "SELECT code, client_id, expires_at, used, redirect_uri, created_at FROM authorization_code WHERE code = $1"

	row := db.QueryRowContext(ctx, query, code)
	var authCode models.AuthorizationCode

	// scan the result to authCode
	err := row.Scan(&authCode.Code, &authCode.ClientID, &authCode.ExpiresAt, &authCode.Used, &authCode.RedirectURI, &authCode.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &authCode, nil
}

func (a AuthorizationCode) InsertAuthorizationCode(ctx context.Context, authCode models.AuthorizationCode, db *sql.Conn) error {

	// query to insert authorization code
	query := "INSERT INTO authorization_codes (code, client_id, expires_at, used, redirect_uri) VALUES ($1, $2, $3, $4, $5)"

	// execute the query
	_, err := db.ExecContext(ctx, query, authCode.Code, authCode.ClientID, authCode.ExpiresAt, authCode.Used, authCode.RedirectURI)
	if err != nil {
		return err
	}
	return nil
}
