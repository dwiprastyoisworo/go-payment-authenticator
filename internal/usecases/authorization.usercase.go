package usecases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/database"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/helpers"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/models"
	"log"
	"slices"
	"time"
)

type Authorization struct {
	db               *sql.Conn
	repositoryClient database.ClientRepository
	repositoryAuth   database.AuthorizationCodeRepository
}

func NewAuthorization(db *sql.Conn, repositoryClient database.ClientRepository, repositoryAuth database.AuthorizationCodeRepository) AuthorizationUsecase {
	return &Authorization{db: db, repositoryClient: repositoryClient, repositoryAuth: repositoryAuth}
}

type AuthorizationUsecase interface {
	RequestAuthorization(ctx context.Context, req models.AuthorizationRequest) (string, error)
}

func (a Authorization) RequestAuthorization(ctx context.Context, req models.AuthorizationRequest) (string, error) {

	// get client by client_id
	clientResponse, err := a.repositoryClient.GetClientById(ctx, req.ClientID, a.db)
	if err != nil {
		log.Println(err)
		return "", errors.New("client not found")
	}

	// validate client redirect uri
	if slices.Contains(clientResponse.RedirectURIs, req.RedirectURI) == false {
		return "", errors.New("redirect uri not allowed")
	}

	code := helpers.GenerateRandomString(32)

	authCode := models.AuthorizationCode{
		Code:        code,
		ClientID:    clientResponse.ID,
		ExpiresAt:   time.Now().Add(10 * time.Minute),
		Used:        false,
		RedirectURI: req.RedirectURI,
	}

	// insert authorization code
	err = a.repositoryAuth.InsertAuthorizationCode(ctx, authCode, a.db)
	if err != nil {
		log.Println(err)
		return "", errors.New("failed to insert authorization code")
	}

	return fmt.Sprintf("%s?code=%s", req.RedirectURI, code), nil

}
