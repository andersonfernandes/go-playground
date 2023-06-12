package dbconnection

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Get() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
