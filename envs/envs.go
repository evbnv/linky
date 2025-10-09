package envs

import (
	"os"
)

var ServerEnvs Envs

type Envs struct {
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_NAME     string
	POSTGRES_USE_SSL  string
}

func LoadEnvs() error {
	ServerEnvs.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	ServerEnvs.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	ServerEnvs.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	ServerEnvs.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	ServerEnvs.POSTGRES_NAME = os.Getenv("POSTGRES_NAME")
	ServerEnvs.POSTGRES_USE_SSL = os.Getenv("POSTGRES_USE_SSL")

	return nil
}
