package drivers

import (
	"database/sql"
	"log"
	"os"
)

func ConnectPostgresDB() *sql.DB {
	foo := os.Getenv("POSTGRES_SQL_URL")
	db, err := sql.Open("postgres", foo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
