package config

import (
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ProductServicePort     string
	DATABASE_URL           string
	GrpcUserServiceUrl     string
	GRPCProductServicePort string
}

func LoadEnvData() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config := &AppConfig{
		ProductServicePort:     os.Getenv("ProductServicePort"),
		DATABASE_URL:           os.Getenv("DATABASE_URL"),
		GrpcUserServiceUrl:     os.Getenv("USER_GRPC_URL"),
		GRPCProductServicePort: os.Getenv("GRPCProductServicePort"),
	}

	return config
}
