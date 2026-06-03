package route

import (
	
	 "web-server/handler"
	"github.com/gin-gonic/gin"
)



func RegisterRoute(router *gin.Engine )  {

	router.POST("/count/:Id", handler.Filecount)



	
}