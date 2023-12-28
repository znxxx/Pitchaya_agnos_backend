package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := "user=root dbname=test_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func GetDB() *sql.DB {
	return db
}
