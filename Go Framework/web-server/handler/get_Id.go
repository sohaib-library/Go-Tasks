package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Getid godoc
// @Summary      Get result by ID
// @Description  Retrieve a single analyzed file record by its ID for the authenticated user
// @Tags         analyzer
// @Produce      json
// @Security     BearerAuth
// @Param        Id  path      int  true  "Record ID"
// @Success      200  {object}  models.File
// @Failure      400  {object}  map[string]string  "Invalid ID"
// @Failure      401  {object}  map[string]string  "Unauthorized"
// @Failure      404  {object}  map[string]string  "Document not found"
// @Router       /result/{Id} [get]
func (getbyid *Handler) Getid(ctx *gin.Context) {
	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)
	if err != nil || Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid int value"})
		return
	}

	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	f, err := getbyid.Analyzer.GetByID(Id, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	ctx.JSON(http.StatusOK, f)
}
