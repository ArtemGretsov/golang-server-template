package errormw

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ArtemGretsov/golang-server-template/src/tools/errorstool"
)

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	requestId := ctx.GetRespHeader(fiber.HeaderXRequestID)

	if httpError, ok := err.(*errorstool.HTTPError); ok {
		responseError := *httpError

		if responseError.StatusCode == fiber.StatusInternalServerError {
			responseError.Message = errorstool.InternalServerErrorMessage
		}

		return ctx.Status(responseError.StatusCode).JSON(errorstool.HTTPErrorRequestID{
			HTTPError: &responseError,
			RequestID: requestId,
		})
	}

	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(errorstool.HTTPErrorRequestID{
			HTTPError: errorstool.NewHTTPInternalServerError(errorstool.InternalServerErrorMessage),
			RequestID: requestId,
		})
}
