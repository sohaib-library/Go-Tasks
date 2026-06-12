package service

import (
	// "go/token"
	// "net/http"
	"web-server/models"
	"web-server/repo"
	"web-server/utils"

	// "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(login models.Login) ( error , string ){

	user, err := repo.Login(login)
	if err != nil {
		return  err , ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(login.PASSWORD))
	if err != nil {
		return  err , ""
	}

	token ,err := utils.GenerateJWT(login.EMAIL)
	if err != nil {
		return  err , ""
	}
	

	return  nil , token
}
