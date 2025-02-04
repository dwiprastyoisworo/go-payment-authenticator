package handlers

import (
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/usecases"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/models"
	"github.com/gofiber/fiber/v3"
)

type Authorization struct {
	authorizationUsecase usecases.AuthorizationUsecase
}

func NewAuthorizationHandler(authorizationUsecase usecases.AuthorizationUsecase) *Authorization {
	return &Authorization{authorizationUsecase: authorizationUsecase}
}

func (r *Authorization) RequestAuthorization(c fiber.Ctx) error {
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")

	if clientID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "client_id is required"})
	}

	if redirectURI == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "redirect_uri is required"})
	}

	req := models.AuthorizationRequest{
		ClientID:    clientID,
		RedirectURI: redirectURI,
	}

	response, err := r.authorizationUsecase.RequestAuthorization(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"url": response})

}
