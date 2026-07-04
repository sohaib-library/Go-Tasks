package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func  (delete *Handler)Deleteid(ctx *gin.Context) {
	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)

	if err != nil || Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The Id does not exist"})
		return
	}

	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	_, err = delete.Analyzer.DeleteByID(Id, userID)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
