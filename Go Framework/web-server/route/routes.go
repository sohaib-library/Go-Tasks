package route

import (
	"web-server/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {

	router.POST("/count", handler.Filecount)
	router.GET("/results", handler.Getdata)
	router.GET("/result/:Id", handler.Getid)
	router.DELETE("/result/:Id", handler.Deleteid)
	router.PATCH("/result/:Id", handler.Updateid)

}
