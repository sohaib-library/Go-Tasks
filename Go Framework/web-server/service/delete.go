package service

import (
	"database/sql"
	"web-server/repo"
)

func DeleteByID(db *sql.DB, id int) (int64, error) {

	return repo.DeleteResultByID(db , id)

}