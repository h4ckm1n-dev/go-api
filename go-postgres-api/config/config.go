package config

import (
	"os"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
}

func LoadConfig() Config {
	return Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://keycloak:password@localhost:5432/keycloak"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8088"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
