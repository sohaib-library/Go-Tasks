package service

import (
	"database/sql"
	"web-server/models"
	"web-server/repo"
)

func GetAll(db *sql.DB) ([]models.File, error) {

	return repo.GetAllResults(db)
	

	
}