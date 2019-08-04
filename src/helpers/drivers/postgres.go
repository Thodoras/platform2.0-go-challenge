package drivers

import (
	"database/sql"
	"log"
	"os"
)

func ConnectPostgresDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_SQL_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
