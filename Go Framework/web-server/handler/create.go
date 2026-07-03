package handler

import (
	"github.com/gin-gonic/gin"
)

func (create *Handler) Filecount(ctx *gin.Context) {
	status, response := create.Analyzer.CreateById(ctx)
	ctx.JSON(status, response)
}
