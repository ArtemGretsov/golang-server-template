package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ArtemGretsov/golang-server-template/internal/middlewares/authmw"
	"github.com/ArtemGretsov/golang-server-template/internal/middlewares/validationmw"
)

func Controller(r fiber.Router) {
  router := r.Group("/auth")
	router.Post(
			"/signup",
			validationmw.ValidateBodyMiddleware(SignupReqDto{}),
			Service.Signup,
	)

	router.Post(
		"/signin",
		validationmw.ValidateBodyMiddleware(SigninReqDto{}),
		Service.Signin,
	)

	router.Get("/user", authmw.Middleware, Service.GetCurrentUser)
}