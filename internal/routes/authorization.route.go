package routes

import (
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/handlers"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/repositories/database"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/usecases"
)

func (r Routes) Authorization() {
	repositoryClient := database.NewClientRepository()
	repositoryAuth := database.NewAuthorizationCodeRepository()
	authorizationUsecase := usecases.NewAuthorization(r.db, repositoryClient, repositoryAuth)
	authorizationHandler := handlers.NewAuthorizationHandler(authorizationUsecase)

	authGroup := r.app.Group("/authorization")
	authGroup.Get("/request", authorizationHandler.RequestAuthorization)

}
