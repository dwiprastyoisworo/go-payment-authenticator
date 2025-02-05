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

func (r *Authorization) RequestToken(c fiber.Ctx) error {
	code := c.FormValue("code")
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "code is required"})
	}

	if clientID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "client_id is required"})
	}

	if clientSecret == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "client_secret is required"})
	}

	req := models.TokenRequest{
		Code:         code,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	response, err := r.authorizationUsecase.RequestToken(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
