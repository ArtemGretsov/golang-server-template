package authmw

import (
	"github.com/gofiber/fiber/v2"
	"regexp"

	"github.com/ArtemGretsov/golang-server-template/internal/database"
	"github.com/ArtemGretsov/golang-server-template/internal/tools/errorstool"
)

type JWTPayload struct {
	ID int
	Login string
	Name string
}

func Middleware(ctx *fiber.Ctx) error {
	rCtx := ctx.UserContext()
	authorizationHeader := ctx.Request().Header.Peek("Authorization")
	token := regexp.
		MustCompile(`Bearer (.+)`).
		ReplaceAllString(string(authorizationHeader), `$1`)

	jwtPayload, err := ParseJWT(token)
	if err != nil {
		return errorstool.NewHTTPError(fiber.StatusUnauthorized, "unauthorized")
	}

	DB := database.DB()

	user, err := DB.User.Get(rCtx, jwtPayload.ID)

	if err != nil || !user.IsActive {
		return errorstool.NewHTTPError(fiber.StatusUnauthorized, "unauthorized")
	}

	ctx.Locals("user", user)

	return ctx.Next()
}

