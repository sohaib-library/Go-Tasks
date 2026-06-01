package handler

import (
	"io"
	"net/http"
	"strconv"
	"web-server/models"
	"web-server/repo"
	wordcount "web-server/service"

	"github.com/gin-gonic/gin"
)

func Filecount(ctx *gin.Context) {

	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Convert "})
		return
	}

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed upload file  "})
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	letters, words, lines, spaces, special := wordcount.Filecontext(Id, string(data))

	// Save result to database
	if err := repo.SaveResult(repo.DB, words, letters, spaces); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save result to DB"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, models.WordCountResponse{
		WordCount:        words,
		LetterCount:      letters,
		LineCount:        lines,
		SpaceCount:       spaces,
		SpecialCharCount: special,
	})
}
