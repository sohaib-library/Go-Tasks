package repo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Database(envpath string) *sql.DB {

	if err := godotenv.Load(envpath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envpath, err)
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Database credentials are not fully set in the environment variables")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	log.Println("Database connected successfully")

	return DB
}

func SaveResult(db *sql.DB, words, letters, spaces, lines, specialchar  int) error {
	query := `INSERT INTO result (total_words, total_letters, total_spaces, total_lines, total_special_char ) VALUES ($1, $2, $3, $4, $5 )`
	_, err := db.Exec(query, words, letters, spaces, lines, specialchar  )
	return err
}
