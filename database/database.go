package database

import (
	"database/sql"
	"fmt"
	"linky/envs"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) SaveURL(shortURL, longURL string) error {
	query := "INSERT INTO urls (short_url, long_url) VALUES ($1, $2)"
	_, err := s.db.Exec(query, shortURL, longURL)
	return err
}

func (s *PostgresStore) GetURL(shortURL string) (string, error) {
	query := "SELECT long_url FROM urls WHERE short_url = $1"
	row := s.db.QueryRow(query, shortURL)
	var longURL string
	if err := row.Scan(&longURL); err != nil {
		return "", err
	}
	return longURL, nil
}

func InitDatabase() (*sql.DB, error) {
	envs := &envs.ServerEnvs
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		envs.POSTGRES_HOST,
		envs.POSTGRES_USER,
		envs.POSTGRES_PASSWORD,
		envs.POSTGRES_NAME,
		envs.POSTGRES_PORT,
		envs.POSTGRES_USE_SSL)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createURLTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createURLTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
    	short_url VARCHAR(255) PRIMARY KEY,
    	long_url TEXT NOT NULL
	)`
	log.Println("Create table if not exists")
	_, err := db.Exec(query)
	return err
}
