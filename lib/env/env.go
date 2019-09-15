package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Getenv is a wrapper around os.Getenv(), nothing more
func Getenv(key string) string {
	return os.Getenv(key)
}

// GetenvWithFallback looks up an env variable and if not set returns a default value
func GetenvWithFallback(key string, fallback string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return v
}

// GetenvBool returns an env variable as a boolean
func GetenvBool(key string) bool {
	v, err := strconv.ParseBool(os.Getenv(key))
	if err == nil {
		return v
	}
	return false
}
