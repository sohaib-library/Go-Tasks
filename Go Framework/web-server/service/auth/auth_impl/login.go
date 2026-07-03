package auth_service

import (
	"web-server/models"
	"web-server/utils"

	"golang.org/x/crypto/bcrypt"
)

func (a *Auth) Login(login models.Login) (string, error) {
	user, err := a.authRepo.Login(login)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(login.PASSWORD))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.ID, login.EMAIL)
	if err != nil {
		return "", err
	}

	return token, nil
}
