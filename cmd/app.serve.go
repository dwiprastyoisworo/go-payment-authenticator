package main

import (
	"context"
	"fmt"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/routes"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/config"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	ctx := context.Background()
	// setup user config
	userConfig, err := config.AppConfigInit()
	if err != nil {
		// handle error
		panic(err)
	}

	// setup postgres connection
	db, err := database.PostgresInit(userConfig, ctx)
	if err != nil {
		// handle error
		panic(err)
	}

	// start fiber http server
	app := fiber.New(
		fiber.Config{
			AppName:       userConfig.App.Name,
			CaseSensitive: true,
		})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("server is running")
	})

	// setup routes
	appRoute := routes.NewRoutes(app, db)
	appRoute.Authorization()

	// start port
	err = app.Listen(fmt.Sprintf(":%d", userConfig.App.Port))

	if err != nil {
		panic(err)
	}
}
