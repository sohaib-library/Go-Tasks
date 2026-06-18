package route

import (
	"web-server/handler"
	"web-server/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {

	//  Authentication

	router.POST("/signup", handler.SignUP)
	router.POST("/login", handler.Login)


	// Access After Validation 

	router.POST("/count", middleware.AuthenticationMiddleware, handler.Filecount)
	router.GET("/results", middleware.AuthenticationMiddleware, handler.Getdata)
	router.GET("/result/:Id",middleware.AuthenticationMiddleware, handler.Getid)
	router.DELETE("/result/:Id",middleware.AuthenticationMiddleware, handler.Deleteid)
	router.PATCH("/result/:Id", middleware.AuthenticationMiddleware,handler.Updateid)

}
