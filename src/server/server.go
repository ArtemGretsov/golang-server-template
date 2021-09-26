package server

import (
	"github.com/ArtemGretsov/golang-server-template/src/middlewares/errormw"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errormw.DefaultErrorHandler,
	})

	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	return app
}

func InitForTest() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errormw.DefaultErrorHandler,
	})

	app.Use(recover.New())

	return app
}
