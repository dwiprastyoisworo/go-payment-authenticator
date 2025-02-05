package usecases

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/database"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/redis"
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
	repositoryRedis  redis.RedisRepository
}

func NewAuthorization(db *sql.Conn, repositoryClient database.ClientRepository, repositoryAuth database.AuthorizationCodeRepository, repositoryRedis redis.RedisRepository) AuthorizationUsecase {
	return &Authorization{db: db, repositoryClient: repositoryClient, repositoryAuth: repositoryAuth, repositoryRedis: repositoryRedis}
}

type AuthorizationUsecase interface {
	RequestAuthorization(ctx context.Context, req models.AuthorizationRequest) (string, error)
}

// RequestAuthorization is a function to request authorization
func (a Authorization) RequestAuthorization(ctx context.Context, req models.AuthorizationRequest) (string, error) {

	var clientResponse *models.Client
	clientResponse = &models.Client{}

	// get client by key on redis
	client, err := a.repositoryRedis.Get(ctx, req.ClientID)

	// if client not found on redis
	if err != nil || client == "" {
		log.Println(fmt.Sprintf("client not found on redis, client_id: %s", req.ClientID))

		// get client by client_id
		clientResponse, err = a.repositoryClient.GetClientById(ctx, req.ClientID, a.db)
		if err != nil {
			log.Println(err)
			return "", errors.New("client not found")
		}

		// json marshal client
		jsonClient, err := json.Marshal(clientResponse)
		if err != nil {
			log.Println(err)
			return "", errors.New("client not found")
		}

		// set client to redis
		err = a.repositoryRedis.Set(ctx, req.ClientID, string(jsonClient), 24*time.Hour)

	} else {
		// json unmarshal client
		err := json.Unmarshal([]byte(client), clientResponse)
		if err != nil {
			fmt.Println("Error:", err)
			return "", errors.New("client not found")
		}
	}

	// validate client redirect uri
	if slices.Contains(clientResponse.RedirectURIs, req.RedirectURI) == false {
		return "", errors.New("redirect uri not allowed")
	}

	// generate authorization code
	code := helpers.GenerateRandomString(32)

	// create authorization code
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
