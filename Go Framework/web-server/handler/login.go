package handler

import (
	"log"
	"net/http"

	"web-server/models"
	"web-server/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(ctx *gin.Context) {

	var login models.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
		log.Print(err.Error())
		return
	}

	_, err := service.Login(login)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfully "})
}
