package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Filecount(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	if userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idStr := ctx.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid int value"})
		return
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed upload file"})
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	result, err := h.Analyzer.CreateById(userID, id, header.Filename, content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
