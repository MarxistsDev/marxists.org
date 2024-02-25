package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv(envFilePath string) error {
	if err := godotenv.Load(envFilePath); err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}

	return nil
}
