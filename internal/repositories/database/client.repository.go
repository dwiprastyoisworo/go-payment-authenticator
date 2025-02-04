package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/models"
	"github.com/lib/pq"
)

type Client struct {
}

func NewClientRepository() ClientRepository {
	return &Client{}
}

type ClientRepository interface {
	GetClientById(ctx context.Context, id string, db *sql.Conn) (*models.Client, error)
}

func (c Client) GetClientById(ctx context.Context, id string, db *sql.Conn) (*models.Client, error) {

	query := "SELECT id, client_id , client_secret ,name ,enabled, redirect_uris FROM clients WHERE client_id = $1"

	row := db.QueryRowContext(ctx, query, id)
	var client models.Client
	err := row.Scan(&client.ID, &client.ClientID, &client.ClientSecret, &client.Name, &client.Enabled, pq.Array(&client.RedirectURIs))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &client, nil
}
