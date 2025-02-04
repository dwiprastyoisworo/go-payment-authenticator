package routes

import (
	"database/sql"
	"github.com/gofiber/fiber/v3"
)

type Routes struct {
	app *fiber.App
	db  *sql.Conn
}

func NewRoutes(app *fiber.App, db *sql.Conn) *Routes {
	return &Routes{app: app, db: db}
}
