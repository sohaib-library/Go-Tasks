package auth_service

import (
	"web-server/database"
	"web-server/models"

	"database/sql"

	"github.com/sirupsen/logrus"
)

func (a *Auth) SignUp(db *sql.DB, users models.Users) (int64, error) {

	affected, err := a.authRepo.SignUP(database.DB, users.NAME, users.EMAIL, users.PASSWORD)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	if affected == 0 {
		return 0, nil
	}

	return affected, nil
}
