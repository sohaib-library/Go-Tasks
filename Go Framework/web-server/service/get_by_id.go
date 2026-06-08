package service

import (
	"database/sql"
	"web-server/models"
	"web-server/repo"
)

func GetByID(db *sql.DB, id int) (models.File, error) {
	return repo.GetResultByID(db, id)
}
