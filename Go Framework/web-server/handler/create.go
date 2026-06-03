package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"web-server/models"
	"web-server/repo"
	wordcount "web-server/service"

	"github.com/gin-gonic/gin"
)

func Filecount(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "GET method is not allowed for file uploads. Please use POST.",
		})
		return
	}

	// get id from user

	// 1. Get string from form-data

	idstr := ctx.PostForm("id")

	Id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid integer value"})
		return
	}

	// File handling
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed upload file  "})
		return
	}

	// readfie

	data, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	letters, words, lines, spaces, special := wordcount.Filecontext(Id, string(data))

	// Save result to database

	if err := repo.SaveResult(repo.DB, words, letters, spaces, lines, special); err != nil {
		fmt.Println("error", err)
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
