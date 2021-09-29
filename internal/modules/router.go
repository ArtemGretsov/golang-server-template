package modules

import (
	"github.com/ArtemGretsov/golang-server-template/internal/modules/auth"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("/api")
	auth.Controller(api)
}
