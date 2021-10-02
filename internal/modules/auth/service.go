package auth

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	userModel "github.com/ArtemGretsov/golang-server-template/internal/database/schemagen/user"

	"github.com/ArtemGretsov/golang-server-template/internal/database"
	"github.com/ArtemGretsov/golang-server-template/internal/database/schemagen"
	"github.com/ArtemGretsov/golang-server-template/internal/middlewares/authmw"
	"github.com/ArtemGretsov/golang-server-template/internal/tools/errorstool"
)

type ServiceType struct {}

var Service = ServiceType{}

const LoginOrPasswordInvalidMessage = "login or password invalid"

func (s *ServiceType) Signin(ctx *fiber.Ctx) error {
	rCtx := ctx.UserContext()
	DB := database.DB()
	body := ctx.Locals("body").(*SigninReqDto)

	userResult, err := DB.User.Query().Where(userModel.Login(body.Login)).First(rCtx)

	if err != nil {
		return errorstool.NewHttpBadRequestError(LoginOrPasswordInvalidMessage)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(body.Password))

	if err != nil {
		return errorstool.NewHttpBadRequestError(LoginOrPasswordInvalidMessage)
	}

	token, err := authmw.CreateJWT(authmw.JWTPayload{
		ID:    userResult.ID,
		Name:  userResult.Name,
		Login: userResult.Login,
	})

	return ctx.Status(fiber.StatusOK).JSON(SignupResDto{
		ID:    userResult.ID,
		Name:  userResult.Name,
		Login: userResult.Login,
		Token: token,
	})
}

func (s *ServiceType) Signup(ctx *fiber.Ctx) error {
	rCtx := ctx.UserContext()
	DB := database.DB()
	body := ctx.Locals("body").(*SignupReqDto)
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return errorstool.NewHTTPInternalServerError(err.Error())
	}

	user, err := DB.User.Create().
		SetName(body.Name).
		SetPassword(string(hashPassword)).
		SetLogin(body.Login).
		Save(rCtx)

	if schemagen.IsConstraintError(err) {
		return errorstool.NewHttpBadRequestError("this login already exists")
	}

	if err != nil {
		return errorstool.NewHTTPInternalServerError(err.Error())
	}

	token, err := authmw.CreateJWT(authmw.JWTPayload{
		ID:    user.ID,
		Name:  user.Name,
		Login: user.Login,
	})

	if err != nil {
		return errorstool.NewHTTPInternalServerError(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(SignupResDto{
		ID:    user.ID,
		Name:  user.Name,
		Login: user.Login,
		Token: token,
	})
}

func (s *ServiceType) GetCurrentUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*schemagen.User)

	return ctx.Status(fiber.StatusOK).JSON(CurrentUserReqDto{
		Login: user.Login,
		Name: user.Name,
		ID: user.ID,
	})
}
