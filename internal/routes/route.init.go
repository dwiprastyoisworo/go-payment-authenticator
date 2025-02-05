package routes

import (
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

type Routes struct {
	app   *fiber.App
	db    *sql.Conn
	redis *redis.Client
}

func NewRoutes(app *fiber.App, db *sql.Conn, redis *redis.Client) *Routes {
	return &Routes{app: app, db: db, redis: redis}
}
