package handler

import (
	"log"
	"net/http"
	"regexp"
	"web-server/database"
	"web-server/models"
	"web-server/service"
	"web-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignUP(ctx *gin.Context) {

	var users models.Users

	var emailRegex = regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signup request"})
		log.Print(err.Error())
		return
	}

	hashedPassword, _ := utils.EncryptPassword(users.PASSWORD)

	users.PASSWORD = hashedPassword

	if !emailRegex.MatchString(users.EMAIL) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Enter Valid Email",
		})
		return
	}

	affected, err := service.SignUP(database.DB, users)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email Already Exists"})
		return
	}

	if affected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User was not inserted"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": " SignUp Successfully"})
}
