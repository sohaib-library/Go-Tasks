package repo

import (
	"database/sql"
	"web-server/models"
)

type AuthRepo interface {
	Login(login models.Login) (*models.Users, error)
	SignUP(db *sql.DB, name string, email string, password string) (int64, error)
}
