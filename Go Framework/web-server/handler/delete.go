package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Deleteid godoc
// @Summary      Delete result by ID
// @Description  Delete an analyzed file record by its ID for the authenticated user
// @Tags         analyzer
// @Produce      json
// @Security     BearerAuth
// @Param        Id  path      int  true  "Record ID"
// @Success      200  {object}  map[string]string  "Deleted successfully"
// @Failure      400  {object}  map[string]string  "Invalid ID"
// @Failure      401  {object}  map[string]string  "Unauthorized"
// @Failure      404  {object}  map[string]string  "Not found"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /result/{Id} [delete]
func (delete *Handler) Deleteid(ctx *gin.Context) {
	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)

	if err != nil || Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The Id does not exist"})
		return
	}

	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	_, err = delete.Analyzer.DeleteByID(Id, userID)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
