package initializers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnv_Successful(t *testing.T) {
	// Arrange
	file, err := os.CreateTemp("", "*.env")

	env := []byte("PROBLEM=\"capitalism\"\nSOLUTION=\"revolution!\"\n")
	file.Write(env)
	file.Close()
	defer os.Remove(file.Name())

	// Act
	err = LoadEnv(file.Name())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "capitalism", os.Getenv("PROBLEM"))
	assert.Equal(t, "revolution!", os.Getenv("SOLUTION"))
}

func TestLoadEnv_Failure_FileDoesntExist(t *testing.T) {
	// Act
	err := LoadEnv("fake.env")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "Error loading .env file: open fake.env: no such file or directory", err.Error())
}
