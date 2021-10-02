package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	"regexp"
	"sync"
)

var once sync.Once
var envTable map[string]string

func Get() map[string]string {
	once.Do(func() {
		envFiles := getEnvFilePaths([]string{".env", ".env.local"})
		envTable, _ = godotenv.Read(envFiles...)
		envTable["POSTGRES_CONNECT_STRING"] = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			envTable["DB_USER"],
			envTable["DB_PASS"],
			envTable["DB_HOST"],
			envTable["DB_PORT"],
			envTable["DB_NAME"],
		)
	})

	return envTable
}

func getEnvFilePaths(envFiles []string) []string {
	testEnv := os.Getenv("TEST_EVN")

	if testEnv == "" {
		return envFiles
	}

	currentWorkDirectory, _ := os.Getwd()

	rootPath := regexp.
		MustCompile(`(.+)/internal/.+`).
		ReplaceAllString(currentWorkDirectory, `$1`)

	var fullPathEnvFiles []string
	for _, v := range envFiles {
		fullPathEnvFiles = append(fullPathEnvFiles, path.Join(rootPath, v))
	}

	return fullPathEnvFiles
}
