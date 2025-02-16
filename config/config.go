package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/joho/godotenv"
)

var projectDirName = "DistributedKeyValueStore"

// GetEnvParam returns the value of the environment variable given by key.
// It will also read from the .env file that is present in the project's root directory.
func GetEnvParam(key string) string {
	reg := regexp.MustCompile("^(.*/" + projectDirName + ")")
	cwd, _ := os.Getwd()
	rootPath := reg.Find([]byte(cwd))

	err := godotenv.Load(filepath.Join(string(rootPath), ".env"))

	if err != nil {
		fmt.Printf("Error loading .env file. " + err.Error())
	}

	return os.Getenv(key)
}
