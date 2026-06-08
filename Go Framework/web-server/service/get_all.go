package service

import (
	"database/sql"
	"web-server/models"
	"web-server/repo"
)

// GetAll returns all word count records from the DB
func GetAll(db *sql.DB) ([]models.File, error) {
	return repo.GetAllResults(db)
}
