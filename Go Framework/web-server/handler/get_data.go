package handler

import (
	"net/http"
	"web-server/database"
	// "web-server/repo"
	"web-server/service"

	"github.com/gin-gonic/gin"
)

func Getdata(ctx *gin.Context) {
	results, err := service.GetAll(database.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}
