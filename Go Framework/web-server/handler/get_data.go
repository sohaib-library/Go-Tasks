package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Getdata godoc
// @Summary      Get all results
// @Description  Retrieve all analyzed file records for the authenticated user
// @Tags         analyzer
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.File
// @Failure      401  {object}  map[string]string  "Unauthorized"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /results [get]
func (getdata *Handler) Getdata(ctx *gin.Context) {
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
