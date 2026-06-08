package database

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migration/*.sql
var embedMigrations embed.FS

func Migertions(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("goose set dialect error: %v", err)
	}

	if err := goose.Up(db, "migration"); err != nil {
		log.Fatalf("goose migration failed: %v", err)
	}

	log.Println("Migrations ran successfully")
}
