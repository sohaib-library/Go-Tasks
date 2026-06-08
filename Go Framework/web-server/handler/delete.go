package handler

import (
	"net/http"
	"strconv"
	"web-server/database"
	"web-server/repo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Deleteid(ctx *gin.Context) {
	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)

	if err != nil || Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The Id does not exist"})
		return
	}

	affected, err := repo.DeleteResultByID(database.DB, Id)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if affected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
