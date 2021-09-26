package validationmw

import (
	"encoding/json"
	"github.com/ArtemGretsov/golang-server-template/src/server"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

	r := httptest.NewRequest("GET", TestRoute + query, nil)

	return app.Test(r)
}

func Test_Validation_ValidateBodyMiddleware_Success(t *testing.T) {
	jsonValidMockData, _ := json.Marshal(MockData{
		Name:     "Timber",
		Login:    "Saw",
		Password: "qwerty",
	})

	response, err := sendRequestWithBody(jsonValidMockData)

	if err != nil {
		t.Fatalf("sending request error: %s", err.Error())
		return
	}

	if response.StatusCode != fiber.StatusOK {
		t.Fatal("success validation error")
		return
	}
}

func Test_Validation_ValidateBodyMiddleware_Fail(t *testing.T) {
	jsonInvalidMockData, _ := json.Marshal(MockData{
		Name:     "Timber",
		Password: "qwerty",
	})

	response, err := sendRequestWithBody(jsonInvalidMockData)

	if err != nil {
		t.Fatalf("Sending request error: %s", err.Error())
		return
	}

	if response.StatusCode != fiber.StatusBadRequest {
		t.Fatal("Failed validation error")
		return
	}
}

func Benchmark_Validation_ValidateBodyMiddleware(b *testing.B) {
	jsonValidMockData, _ := json.Marshal(MockData{
		Name:     "Timber",
		Login:    "Saw",
		Password: "qwerty",
	})

	for i := 0; i < b.N; i++ {
		_, _ = sendRequestWithBody(jsonValidMockData)
	}
}

func Benchmark_Validation_ValidateQueryMiddleware(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sendRequestWithQuery("name=Timber&login=saw&password=query")
	}
}

func Test_Validation_ValidateQueryMiddleware_Success(t *testing.T) {
	response, err := sendRequestWithQuery("?name=Timber&login=saw&password=query")

	if err != nil {
		t.Fatalf("sending request error: %s", err.Error())
		return
	}

	if response.StatusCode != fiber.StatusOK {
		t.Fatal("success validation error")
		return
	}
}

func Test_Validation_ValidateQueryMiddleware_Fail(t *testing.T) {
	response, err := sendRequestWithQuery("?name=Timber&password=query")

	if err != nil {
		t.Fatalf("sending request error: %s", err.Error())
		return
	}

	if response.StatusCode != fiber.StatusBadRequest {
		t.Fatal("success validation error")
		return
	}
}

