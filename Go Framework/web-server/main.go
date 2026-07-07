// @title           Web Server API
// @version         1.0
// @description     REST API built with Gin
// @host            localhost:8000
// @BasePath        /

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Enter your JWT token as: Bearer <token>

package main

import (
	"web-server/database"
	_ "web-server/docs"
	"web-server/route"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	DB := database.Database(".env")
	defer DB.Close()

	database.Migertions(DB)

	router := gin.Default()


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
 
	route.RegisterRoute(router, DB)

	
	router.Run(":8000")
}
