package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("NOTES_DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
