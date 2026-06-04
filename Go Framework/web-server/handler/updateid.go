package handler

import (
	"net/http"
	"strconv"
	"web-server/models"
	"web-server/repo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Updateid(ctx *gin.Context) {

	// Get Specific Id from user
	key := ctx.Param("Id")
	Id, _ := strconv.Atoi(key)

	if Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing Id"})
		return
	}

	var final models.File

	if err := ctx.ShouldBindJSON(&final); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := repo.DB.Exec(`
		UPDATE result
		SET total_words       = $1,
		    total_letters     = $2,
		    total_spaces      = $3,
		    total_lines       = $4,
		    total_special_char = $5
		WHERE id = $6
	`, final.TotalWords, final.TotalLetters, final.TotalSpaces, final.TotalLines, final.TotalSpecial, Id)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
