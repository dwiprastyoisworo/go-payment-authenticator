package main

import (
	"context"
	"fmt"
	"github.com/dwiprastyoisworo/go-payment-authenticator/internal/routes"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/config"
	"github.com/gofiber/fiber/v3"
)

func main() {
	ctx := context.Background()
	// setup user config
	cfg, err := config.AppConfigInit()
	if err != nil {
		// handle error
		panic(err)
	}

	// setup postgres connection
	db, err := cfg.PostgresInit(ctx)
	if err != nil {
		// handle error
		panic(err)
	}

	// setup redis connection
	redisClient := cfg.RedisInit()

	// start fiber http server
	app := fiber.New(
		fiber.Config{
			AppName:       cfg.App.Name,
			CaseSensitive: true,
		})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("server is running")
	})

	// setup routes
	appRoute := routes.NewRoutes(app, db, redisClient, cfg)
	appRoute.Authorization()

	// start port
	err = app.Listen(fmt.Sprintf(":%d", cfg.App.Port))

	if err != nil {
		panic(err)
	}
}
