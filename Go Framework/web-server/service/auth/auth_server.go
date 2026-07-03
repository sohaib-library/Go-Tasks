package auth

import (
	"database/sql"
	"web-server/models"
)

type AuthServices interface {
	Login(login models.Login) (string, error)
	SignUp(db *sql.DB, users models.Users) (int64, error)
}
