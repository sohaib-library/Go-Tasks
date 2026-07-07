package handler

import (
	"net/http"
	"strconv"
	"web-server/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Updateid godoc
// @Summary      Update result by ID
// @Description  Update an analyzed file record by its ID for the authenticated user
// @Tags         analyzer
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        Id    path      int          true  "Record ID"
// @Param        file  body      models.File  true  "Updated file data"
// @Success      200   {object}  models.File
// @Failure      400   {object}  map[string]string  "Invalid ID or body"
// @Failure      401   {object}  map[string]string  "Unauthorized"
// @Failure      404   {object}  map[string]string  "Record not found"
// @Failure      500   {object}  map[string]string  "Internal server error"
// @Router       /result/{Id} [patch]
func (update *Handler) Updateid(ctx *gin.Context) {

	key := ctx.Param("Id")
	Id, err := strconv.Atoi(key)
	if err != nil || Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing Id"})
		return
	}

	var f models.File
	if err := ctx.ShouldBindJSON(&f); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	affected, err := update.Analyzer.UpdateByID(Id, userID, f)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	if affected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	f.ID = Id
	ctx.JSON(http.StatusOK, f)
}
