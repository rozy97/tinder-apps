package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	AppPort string
}

func InitEnvironment() Environment {
	godotenv.Load()

	return Environment{
		AppPort: os.Getenv("APP_PORT"),
	}
}