package auth

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/ArtemGretsov/golang-server-template/src/database/repositories/userrep"
	"github.com/ArtemGretsov/golang-server-template/src/middlewares/authmw"
	"github.com/ArtemGretsov/golang-server-template/src/tools/errorstool"
)

type ServiceType struct {
	UserRepository userrep.RepositoryInterface
}

var Service = ServiceType{
	UserRepository: userrep.Repository,
}

func (s *ServiceType) LoginUser(login string, password string) userrep.User {
	user, err := s.UserRepository.GetUserByLogin(login)

	if err != nil {
		panic(err)
	}

	return user
}

func (s *ServiceType) Signup(ctx *fiber.Ctx) error {
	body := ctx.Locals("body").(*SignupReqDto)
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return errorstool.NewHTTPInternalServerError(err.Error())
	}

	user, err := s.UserRepository.SaveUser(userrep.User{
		Login: body.Login,
		Password: string(hashPassword),
		Name: body.Name,
	})

	if err != nil {
		return errorstool.NewHTTPInternalServerError(err.Error())
	}

	token, err := authmw.CreateJWT(authmw.JWTPayload{
		ID: int(user.ID),
		Name: user.Name,
		Login: user.Login,
	})

	if err != nil {
		return errorstool.NewHTTPInternalServerError(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
