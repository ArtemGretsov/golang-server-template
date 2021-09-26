package config

import (
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
		MustCompile(`(.+)/src/(.+)`).
		ReplaceAllString(currentWorkDirectory, `$1`)

	var fullPathEnvFiles []string
	for _, v := range envFiles {
		fullPathEnvFiles = append(fullPathEnvFiles, path.Join(rootPath, v))
	}

	return fullPathEnvFiles
}
