package errormw

import (
	"github.com/ArtemGretsov/golang-server-template/src/tools/errorsutil"
	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	requestId := ctx.GetRespHeader(fiber.HeaderXRequestID)

	if httpError, ok := err.(*errorsutil.HTTPError); ok {
		responseError := *httpError

		if responseError.StatusCode == fiber.StatusInternalServerError {
			responseError.Message = errorsutil.InternalServerErrorMessage
		}

		return ctx.Status(responseError.StatusCode).JSON(errorsutil.HTTPErrorRequestID{
			HTTPError: &responseError,
			RequestID: requestId,
		})
	}

	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(errorsutil.HTTPErrorRequestID{
			HTTPError: errorsutil.NewHTTPInternalServerError(errorsutil.InternalServerErrorMessage),
			RequestID: requestId,
		})
}
