package main

import (
	// "web-server/handler"
	"web-server/migration"
	"web-server/repo"
	"web-server/route"

	"github.com/gin-gonic/gin"
)

func main() {

	DB := repo.Database(".env")
	defer DB.Close()
	migration.Migertions(DB)

	router := gin.Default()
	route.RegisterRoute(router)
	router.Run(":8000")
}

//post user file or go route
//get by id
//get all
//delete
