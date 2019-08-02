package drivers

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func ConnectPostgresDB() {
	foo := os.Getenv("POSTGRES_SQL_URL")
	db, err := sql.Open("postgres", foo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
