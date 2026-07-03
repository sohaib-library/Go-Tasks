package auth_service

import (
	"database/sql"
	"web-server/models"
	// "web-server/repo/auth/auth_impl"
)

func (a *Auth) SignUp(db *sql.DB, users models.Users) (int64, error) {

	return a.authRepo.SignUP(db, users.NAME, users.EMAIL, users.PASSWORD)
}
