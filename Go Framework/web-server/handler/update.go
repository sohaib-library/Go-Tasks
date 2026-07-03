package handler

import (
	"net/http"
	"strconv"
	"web-server/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (update *Handler) Updateid(ctx *gin.Context) {

	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)
	if err != nil || Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing Id"})
		return
	}

	var f models.File
	if err := ctx.ShouldBindJSON(&f); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	affected, err := update.Analyzer.UpdateByID(Id, userID, f)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	if affected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	f.ID = Id
	ctx.JSON(http.StatusOK, f)
}
