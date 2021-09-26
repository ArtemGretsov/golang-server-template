package authmw

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"

	"github.com/ArtemGretsov/golang-server-template/src/server"
)

const TestRoute = "/test"

func Test_Middleware_Fail(t *testing.T) {
	app := server.InitForTest()

	app.Get(TestRoute, Middleware, func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	r := httptest.NewRequest("GET", TestRoute, nil)
	r.Header.Set("Authorization", "test_token_jwt")

	response, err := app.Test(r)

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, fiber.StatusUnauthorized, "status code is not unauthorized")
}

func Test_Middleware_Success(t *testing.T) {
	app := server.InitForTest()

	app.Get(TestRoute, Middleware, func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	token, err := CreateJWT(JWTPayload{
		ID: 1,
		Login: "login",
		Name: "name",
	})
	assert.Nil(t, err)

	r := httptest.NewRequest("GET", TestRoute, nil)
	r.Header.Set("Authorization", "Bearer " + token)

	response, err := app.Test(r)

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, fiber.StatusOK, "status code is not OK")
}