package database

import (
	"database/sql"
	"fmt"
	"linky/envs"
	"log"

	_ "github.com/lib/pq"
)

var PostgresClient *sql.DB

func InitDatabase() error {
	envs := &envs.ServerEnvs
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		envs.POSTGRES_USER,
		envs.POSTGRES_PASSWORD,
		envs.POSTGRES_NAME,
		envs.POSTGRES_USE_SSL)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	} else {
		PostgresClient = db
	}

	if err := PostgresClient.Ping(); err != nil {
		return err
	}

	if err := createURLTable(); err != nil {
		return err
	}

	return nil
}

func createURLTable() error {
	query := `CREATE TABLE IF NOT EXISTS urls (
    	short_url VARCHAR(255) PRIMARY KEY,
    	long_url TEXT NOT NULL
	)`
	log.Println("Create table if not exists")
	_, err := PostgresClient.Exec(query)
	return err
}
