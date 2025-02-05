package routes

import (
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/handlers"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/database"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/redis"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/usecases"
)

func (r Routes) Authorization() {
	repositoryClient := database.NewClientRepository()
	repositoryAuth := database.NewAuthorizationCodeRepository()
	respsitoryAccessToken := database.NewAccessToken()
	repositoryRedis := redis.NewRedis(r.redis)
	authorizationUsecase := usecases.NewAuthorization(r.db, repositoryClient, repositoryAuth, respsitoryAccessToken, repositoryRedis, r.cfg)
	authorizationHandler := handlers.NewAuthorizationHandler(authorizationUsecase)

	authGroup := r.app.Group("/authorization")
	authGroup.Get("/request", authorizationHandler.RequestAuthorization)
	authGroup.Post("/token", authorizationHandler.RequestToken)
}
