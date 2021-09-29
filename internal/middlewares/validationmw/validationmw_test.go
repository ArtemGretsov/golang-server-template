package validationmw

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ArtemGretsov/golang-server-template/internal/server"
)

const TestRoute = "/test"

type MockData struct {
	Name     string `json:"name"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password"`
}

func sendRequestWithBody(json []byte) (*http.Response, error) {
	app := server.InitForTest()

	app.Post(TestRoute, ValidateBodyMiddleware(MockData{}), func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	r := httptest.NewRequest("POST", TestRoute, strings.NewReader(string(json)))
	r.Header.Set("Content-Type", "application/json")

	return app.Test(r)
}

func sendRequestWithQuery(query string) (*http.Response, error) {
	app := server.InitForTest()

	app.Get(TestRoute, ValidateQueryMiddleware(MockData{}), func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	r := httptest.NewRequest("GET", TestRoute+query, nil)

	return app.Test(r)
}

func Test_ValidationMW_ValidateBodyMiddleware_Success(t *testing.T) {
	jsonValidMockData, _ := json.Marshal(MockData{
		Name:     "Timber",
		Login:    "Saw",
		Password: "qwerty",
	})

	response, err := sendRequestWithBody(jsonValidMockData)

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, fiber.StatusOK, "success validation error")
}

func Test_ValidationMW_ValidateBodyMiddleware_Fail(t *testing.T) {
	jsonInvalidMockData, _ := json.Marshal(MockData{
		Name:     "Timber",
		Password: "qwerty",
	})

	response, err := sendRequestWithBody(jsonInvalidMockData)

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, fiber.StatusBadRequest, "fail validation error")
}

func Benchmark_ValidationMW_ValidateBodyMiddleware(b *testing.B) {
	jsonValidMockData, _ := json.Marshal(MockData{
		Name:     "Timber",
		Login:    "Saw",
		Password: "qwerty",
	})

	for i := 0; i < b.N; i++ {
		_, _ = sendRequestWithBody(jsonValidMockData)
	}
}

func Benchmark_ValidationMW_ValidateQueryMiddleware(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sendRequestWithQuery("name=Timber&login=saw&password=query")
	}
}

func Test_ValidationMW_ValidateQueryMiddleware_Success(t *testing.T) {
	response, err := sendRequestWithQuery("?name=Timber&login=saw&password=query")

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, fiber.StatusOK, "success validation error")
}

func Test_ValidationMW_ValidateQueryMiddleware_Fail(t *testing.T) {
	response, err := sendRequestWithQuery("?name=Timber&password=query")

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, fiber.StatusBadRequest, "fail validation error")
}
