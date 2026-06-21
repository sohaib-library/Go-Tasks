package service

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"web-server/models"
	"web-server/repo"

	"github.com/gin-gonic/gin"
)

func CreateById(db *sql.DB, ctx *gin.Context) (int, any) {
	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		return http.StatusUnauthorized, gin.H{"error": "unauthorized"}
	}

	idStr := ctx.PostForm("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return http.StatusBadRequest, gin.H{"error": "Invalid int value"}
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		return http.StatusBadRequest, gin.H{"error": "Failed upload file"}
	}
	defer file.Close()

	if !strings.HasSuffix(header.Filename, ".txt") {
		return http.StatusBadRequest, gin.H{
			"error": "only .txt files allowed"}

	}

	data, err := io.ReadAll(file)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "failed to read file"}
	}

	letters, words, lines, spaces, special := Filecontext(id, string(data))

	err = repo.InsertData(db, userID, words, letters, spaces, lines, special)
	if err != nil {
		log.Print("error : Internal Server Error")
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	return http.StatusOK, &models.WordCountResponse{
		WordCount:        words,
		LetterCount:      letters,
		LineCount:        lines,
		SpaceCount:       spaces,
		SpecialCharCount: special,
	}

}
