package service

import (
	"database/sql"
	"web-server/models"
	"web-server/repo"
)

func UpdateByID(db *sql.DB, id int, f models.File) (int64, error) {
	return repo.UpdateResultByID(db, id, f)
}
