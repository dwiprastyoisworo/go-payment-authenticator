package usecases

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/database"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/redis"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/config"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/helpers"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/models"
	"log"
	"slices"
	"time"
)

type Authorization struct {
	db                    *sql.Conn
	repositoryClient      database.ClientRepository
	repositoryAuth        database.AuthorizationCodeRepository
	repositoryAccessToken database.AccessTokenRepository
	repositoryRedis       redis.RedisRepository
	cfg                   *config.AppConfig
}

func NewAuthorization(db *sql.Conn, repositoryClient database.ClientRepository, repositoryAuth database.AuthorizationCodeRepository, repositoryAccessToken database.AccessTokenRepository, repositoryRedis redis.RedisRepository, cfg *config.AppConfig) AuthorizationUsecase {
	return &Authorization{db: db, repositoryClient: repositoryClient, repositoryAuth: repositoryAuth, repositoryAccessToken: repositoryAccessToken, repositoryRedis: repositoryRedis, cfg: cfg}
}

type AuthorizationUsecase interface {
	RequestAuthorization(ctx context.Context, req models.AuthorizationRequest) (string, error)
	RequestToken(ctx context.Context, req models.TokenRequest) (models.TokenResponse, error)
}

// RequestAuthorization is a function to request authorization
func (a Authorization) RequestAuthorization(ctx context.Context, req models.AuthorizationRequest) (string, error) {

	clientResponse, err := a.getClientWithCaching(ctx, req.ClientID)

	if err != nil {
		return "", err
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

func (a Authorization) RequestToken(ctx context.Context, req models.TokenRequest) (models.TokenResponse, error) {

	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		log.Println(err)
		return models.TokenResponse{}, errors.New("internal server error")
	}
	// get authorization code
	authCode, err := a.repositoryAuth.GetAuthorizationCode(ctx, req.Code, a.db)
	if err != nil {
		log.Println(err)
		return models.TokenResponse{}, errors.New("authorization code not found")
	}

	// validate authorization code
	if authCode.Used == true {
		return models.TokenResponse{}, errors.New("authorization code already used")
	}

	// get client
	clientResponse, err := a.getClientWithCaching(ctx, req.ClientID)
	if err != nil {
		return models.TokenResponse{}, err
	}

	// validate client
	if clientResponse.ClientSecret != req.ClientSecret {
		return models.TokenResponse{}, errors.New("client secret not match")
	}

	tokenJwt := models.Claim{
		ClientID: clientResponse.ClientID,
	}

	expiredToken := time.Now().Add(1 * time.Hour)
	token, err := tokenJwt.GenerateToken(a.cfg.Jwt.SecretKey, expiredToken)
	if err != nil {
		return models.TokenResponse{}, errors.New("failed to generate token")
	}

	// generate access token
	accessToken := models.AccessToken{
		Token:     token,
		ClientID:  clientResponse.ID,
		ExpiresAt: expiredToken,
		Revoked:   false,
	}

	// insert access token
	err = a.repositoryAccessToken.InsertAccessToken(ctx, accessToken, tx)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return models.TokenResponse{}, errors.New("failed to insert access token")
	}

	// update authorization code
	authCode.Used = true
	err = a.repositoryAuth.UpdateAuthorizationCode(ctx, *authCode, tx)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return models.TokenResponse{}, errors.New("failed to update authorization code")
	}
	tx.Commit()
	return models.TokenResponse{
		AccessToken: token,
		ExpiresAt:   accessToken.ExpiresAt,
	}, nil
}

func (a Authorization) getClientWithCaching(ctx context.Context, clientId string) (*models.Client, error) {
	var clientResponse *models.Client
	clientResponse = &models.Client{}

	// get client by key on redis
	client, err := a.repositoryRedis.Get(ctx, clientId)

	// if client not found on redis
	if err != nil || client == "" {
		log.Println(fmt.Sprintf("client not found on redis, client_id: %s", clientId))

		// get client by client_id
		clientResponse, err = a.repositoryClient.GetClientById(ctx, clientId, a.db)
		if err != nil {
			log.Println(err)
			return clientResponse, errors.New("client not found")
		}

		// json marshal client
		jsonClient, err := json.Marshal(clientResponse)
		if err != nil {
			log.Println(err)
			return clientResponse, errors.New("client not found")
		}

		// set client to redis
		err = a.repositoryRedis.Set(ctx, clientId, string(jsonClient), 24*time.Hour)

	} else {
		// json unmarshal client
		err := json.Unmarshal([]byte(client), clientResponse)
		if err != nil {
			fmt.Println("Error:", err)
			return clientResponse, errors.New("client not found")
		}
	}

	return clientResponse, nil
}
