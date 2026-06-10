package handler

import (
	"log"
	"net/http"
	"web-server/database"
	"web-server/models"
	"web-server/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignUP(ctx *gin.Context) {


	var users models.Users

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signup request"})
		log.Print(err.Error())
		return
	}

	affected, err := service.SignUP(database.DB, users)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if affected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User was not inserted"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
