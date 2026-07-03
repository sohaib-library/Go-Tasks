package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (getdata *Handler)Getdata(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	results, err := getdata.Analyzer.GetAll(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}
