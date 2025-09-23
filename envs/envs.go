package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ServerEnvs Envs

type Envs struct {
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_NAME     string
	POSTGRES_USE_SSL  string
}

func LoadEnvs() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ServerEnvs.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	ServerEnvs.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	ServerEnvs.POSTGRES_NAME = os.Getenv("POSTGRES_NAME")
	ServerEnvs.POSTGRES_USE_SSL = os.Getenv("POSTGRES_USE_SSL")

	return nil
}
