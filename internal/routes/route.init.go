package routes

import (
	"database/sql"
	"github.com/dwiprastyoisworo/go-payment-authenticator/lib/config"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

type Routes struct {
	app   *fiber.App
	db    *sql.Conn
	redis *redis.Client
	cfg   *config.AppConfig
}

func NewRoutes(app *fiber.App, db *sql.Conn, redis *redis.Client, cfg *config.AppConfig) *Routes {
	return &Routes{app: app, db: db, redis: redis, cfg: cfg}
}
