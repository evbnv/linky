package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var PostgresClient *sql.DB

func InitDatabase() error {
	connStr := "user=postgres password=gtyIEmdxcfu3783!_12 dbname=linky sslmode=disable"
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
	quary := `CREATE TABLE IF NOT EXISTS urls (
    	short_url VARCHAR(255) PRIMARY KEY,
    	long_url TEXT NOT NULL
	)`
	log.Println("Create table if not exists")
	_, err := PostgresClient.Exec(quary)
	return err
}
