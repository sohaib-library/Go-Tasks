package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (getbyid *Handler) Getid(ctx *gin.Context) {
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

	f, err := getbyid.Analyzer.GetByID(Id, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	ctx.JSON(http.StatusOK, f)
}
