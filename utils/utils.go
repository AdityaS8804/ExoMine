package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
    // Load .env file from the root directory
    err := godotenv.Load()
    if err != nil {
        return fmt.Errorf("error loading .env file: %w", err)
    }
    return nil
}

func GetAPIKey() string {
    return os.Getenv("API_KEY")
}