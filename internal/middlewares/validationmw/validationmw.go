package validationmw

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"

	"github.com/ArtemGretsov/golang-server-template/internal/tools/errorstool"
)

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func validateStruct(schema interface{}) []ErrorResponse {
	var errors []ErrorResponse
	validate := validator.New()
	err := validate.Struct(schema)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func ValidateBodyMiddleware(schema interface{}) func(c *fiber.Ctx) error {
	schemaType := reflect.TypeOf(schema)

	return func(c *fiber.Ctx) error {
		body := reflect.New(schemaType).Interface()

		if err := c.BodyParser(body); err != nil {
			return errorstool.NewHttpBadRequestError("body parsing error")
		}

		errors := validateStruct(body)


		if len(errors) != 0 {
			return errorstool.NewHTTPValidationError(errors)
		}

		c.Locals("body", body)

		return c.Next()
	}
}

func ValidateQueryMiddleware(schema interface{}) func(c *fiber.Ctx) error {
	schemaType := reflect.TypeOf(schema)

	return func(c *fiber.Ctx) error {
		query := reflect.New(schemaType).Interface()

		if err := c.QueryParser(query); err != nil {
			return errorstool.NewHttpBadRequestError("query params parsing error")
		}

		errors := validateStruct(query)

		if len(errors) != 0 {
			return errorstool.NewHTTPValidationError(errors)
		}

		c.Locals("query", query)

		return c.Next()
	}
}