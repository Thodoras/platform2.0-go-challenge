package drivers

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func ConnectPostgresDB() {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_SQL_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
