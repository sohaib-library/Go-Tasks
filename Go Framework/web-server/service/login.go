package service

import (
	"web-server/models"
	"web-server/repo"

	"golang.org/x/crypto/bcrypt"
)

func Login(login models.Login) (*models.Users, error) {

	user, err := repo.Login(login)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(login.PASSWORD))
	if err != nil {
		return nil, err
	}

	return user, nil
}
