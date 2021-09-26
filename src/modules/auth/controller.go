package auth

import (
	"github.com/ArtemGretsov/golang-server-template/src/middlewares/authmw"
	"github.com/ArtemGretsov/golang-server-template/src/middlewares/validationmw"
	"github.com/gofiber/fiber/v2"
)

func Controller(r fiber.Router) {
  router := r.Group("/auth")
	router.Post(
			"/signup",
			authmw.Middleware,
			validationmw.ValidateBodyMiddleware(SignupReqDto{}),
			Service.Signup,
	)
}