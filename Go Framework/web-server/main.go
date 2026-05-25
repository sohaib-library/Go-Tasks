package main

import (
	"web-server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/count/:Id", handler.Filecount)

	router.Run(":8000")
}
