package service

import (
	"database/sql"
	"web-server/models"
	"web-server/repo"
)

func SignUP(db *sql.DB, users models.Users) (int64, error) {
	return repo.SignUP( db, users.NAME, users.EMAIL, users.PASSWORD )
}
