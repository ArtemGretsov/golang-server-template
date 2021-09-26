package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ArtemGretsov/golang-server-template/src/middlewares/validationmw"
)

func Controller(r fiber.Router) {
  router := r.Group("/auth")
	router.Post(
			"/signup",
			validationmw.ValidateBodyMiddleware(SignupReqDto{}),
			Service.Signup,
	)
}