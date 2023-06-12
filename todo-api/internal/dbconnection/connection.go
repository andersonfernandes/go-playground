package dbconnection

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Get() *sql.DB {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

  return db
}
