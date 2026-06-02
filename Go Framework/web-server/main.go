package main

import (
	"web-server/handler"
	"web-server/migration"
	"web-server/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()

	
	DB := repo.Database(".env")
	defer DB.Close() 

	migration.Migertions(DB)

	router.GET("/count/:Id", handler.Filecount)

	router.Run(":8000")
}

//post user file or go route
//get by id 
//get all
//delete