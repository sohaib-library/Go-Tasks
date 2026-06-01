package migration

import (
	"database/sql"
	"log"
)

func Migertions(db *sql.DB) {

	query := `
        CREATE TABLE IF NOT EXISTS files (
        id      SERIAL PRIMARY KEY,
        name    TEXT NOT NULL,
        content TEXT
    );`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations ran successfully")

}
