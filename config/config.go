package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .evn file found, using system environment variables")
	}
}

	func GetEnv(key string, defaultValue string) string {
		value, exists := os.LookupEnv(key)
		if !exists {
			return defaultValue
		}
		return value
}
