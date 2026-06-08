package handler

import (
	"net/http"
	"web-server/database"
	"web-server/repo"

	"github.com/gin-gonic/gin"
)

func Getdata(ctx *gin.Context) {
	results, err := repo.GetAllResults(database.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}
