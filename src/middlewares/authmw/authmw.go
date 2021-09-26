package authmw

import (
	"github.com/gofiber/fiber/v2"
)

type JWTPayload struct {
	ID int
	Login string
	Name string
}

func Middleware(ctx *fiber.Ctx) error {
	return ctx.Next()
}

