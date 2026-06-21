package handler

import (
	"net/http"
	"strconv"
	"web-server/database"
	// "web-server/repo"
	"web-server/service"

	"github.com/gin-gonic/gin"
)

func Getid(ctx *gin.Context) {
	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)
	if err != nil || Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid int value"})
		return
	}

	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	f, err := service.GetByID(database.DB, Id, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	ctx.JSON(http.StatusOK, f)
}
