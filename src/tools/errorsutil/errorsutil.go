package errorsutil

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type HTTPError struct {
	Message     string      `json:"message"`
	MessageCode *string     `json:"messageCode"`
	StatusCode  int         `json:"code"`
	Timestamp   time.Time   `json:"timestamp"`
	Details     interface{} `json:"details"`
}

type HTTPErrorRequestID struct {
	*HTTPError
	RequestID string `json:"requestId"`
}

const InternalServerErrorMessage = "Internal Server Error"

func (h *HTTPError) Error() string {
	return h.Message
}

func NewHTTPError(code int, message string, messages ...string) *HTTPError {
	var messageCode *string

	if len(messages) > 0 {
		messageCode = &messages[0]
	}

	return &HTTPError{
		Message:     message,
		MessageCode: messageCode,
		StatusCode:  code,
		Timestamp:   time.Now(),
	}
}

func NewHTTPNotFoundError(messages ...string) *HTTPError {
	var messageCode *string
	message := "Not Found"

	if len(messages) > 0 {
		message = messages[0]
	}

	if len(messages) > 1 {
		messageCode = &messages[1]
	}

	return &HTTPError{
		Message:     message,
		MessageCode: messageCode,
		StatusCode:  fiber.StatusNotFound,
		Timestamp:   time.Now().UTC(),
	}
}

func NewHttpBadRequestError(message string, messages ...string) *HTTPError {
	var messageCode *string

	if len(messages) > 0 {
		messageCode = &messages[0]
	}

	return &HTTPError{
		Message:     message,
		MessageCode: messageCode,
		StatusCode:  fiber.StatusBadRequest,
		Timestamp:   time.Now().UTC(),
	}
}

func NewHTTPInternalServerError(message string, messages ...string) *HTTPError {
	var messageCode *string

	if len(messages) > 0 {
		messageCode = &messages[0]
	}

	return &HTTPError{
		Message:     message,
		MessageCode: messageCode,
		StatusCode:  fiber.StatusInternalServerError,
		Timestamp:   time.Now().UTC(),
	}
}

func NewHTTPValidationError(details interface{}) *HTTPError {
	return &HTTPError{
		Message:    "Validation Error",
		StatusCode: fiber.StatusBadRequest,
		Timestamp:  time.Now().UTC(),
		Details:    details,
	}
}
