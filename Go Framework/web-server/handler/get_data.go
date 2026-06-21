package handler

import (
	"net/http"
	"web-server/database"
	// "web-server/repo"
	"web-server/service"

	"github.com/gin-gonic/gin"
)

func Getdata(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	results, err := service.GetAll(database.DB, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}
