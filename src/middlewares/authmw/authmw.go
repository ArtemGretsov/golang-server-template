package authmw

import (
	"github.com/gofiber/fiber/v2"
	"regexp"

	"github.com/ArtemGretsov/golang-server-template/src/tools/errorstool"
)

type JWTPayload struct {
	ID int
	Login string
	Name string
}

func Middleware(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Request().Header.Peek("Authorization")
	token := regexp.
		MustCompile(`Bearer (.+)`).
		ReplaceAllString(string(authorizationHeader), `$1`)

	jwtPayload, err := ParseJWT(token)
	if err != nil {
		return errorstool.NewHTTPError(fiber.StatusUnauthorized, "unauthorized")
	}

	ctx.Locals("user", &jwtPayload)

	return ctx.Next()
}

