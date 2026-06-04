package handler

import (
	// "log"
	// "net/http"

	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	// "web-server/models"
	"web-server/repo"

	"github.com/gin-gonic/gin"
)

func Deleteid(ctx *gin.Context) {

	// Get Specific Id from user

	key := ctx.Param("Id")
	Id, _ := strconv.Atoi(key)

	if Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The Id does not exist"})
		return
	}

	row, err := repo.DB.Exec(`

	DELETE FROM result WHERE id = $1
    `, Id)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil || rowsAffected == 0 {

		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
