package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AuthServicePort     string
	SecretKey 			string
	GRPCAuthServicePort string
}

func LoadEnvData() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config := &AppConfig{
		AuthServicePort: os.Getenv("AuthServicePort"),
		SecretKey: 	os.Getenv("SECRET_KEY"),
		GRPCAuthServicePort: os.Getenv("GRPCAuthServicePort"),
	}

	return config
}