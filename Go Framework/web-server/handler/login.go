package handler

import (
	// "database/sql"
	"log"
	"net/http"

	// "web-server/database"
	"web-server/models"
	"web-server/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(ctx *gin.Context)  {

	var login models.Login
	  
     if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signup request"})
		log.Print(err.Error())
		return
	}

	affected, err := service.Login( login)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if affected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User was not inserted"})
		return
	}


	


	
}