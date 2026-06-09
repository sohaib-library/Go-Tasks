package handler

import (
	"web-server/database"
	"web-server/service"

	"github.com/gin-gonic/gin"
)

func Filecount(ctx *gin.Context) {
	status, response := service.CreateById(database.DB, ctx)
	ctx.JSON(status, response)
}
