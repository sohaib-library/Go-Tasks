package service

import (
	"database/sql"
	"web-server/models"
	"web-server/repo"
)

func GetAll(db *sql.DB, userID int) ([]models.File, error) {

	return repo.GetAllResults(db, userID)

}
