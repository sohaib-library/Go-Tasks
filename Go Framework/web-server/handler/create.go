package handler

import (
	"io"
	"net/http"
	"strconv"
	"web-server/database"
	"web-server/service"

	"github.com/gin-gonic/gin"
)

func Filecount(ctx *gin.Context) {


	idstr := ctx.PostForm("id")
	workers, err := strconv.Atoi(idstr)
	if err != nil || workers <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid integer value for 'id'"})
		return
	}


	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}


	result, err := service.CreateWordCount(database.DB, workers, string(data))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save result to DB"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
