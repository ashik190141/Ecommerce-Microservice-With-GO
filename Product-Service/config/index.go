package config

import (
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ProductServicePort string
	DATABASE_URL       string
}

func LoadEnvData() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config := &AppConfig{
		ProductServicePort: os.Getenv("AuthServicePort"),
		DATABASE_URL:       os.Getenv("DATABASE_URL"),
	}

	return config
}
