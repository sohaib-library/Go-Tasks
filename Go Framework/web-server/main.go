package main

import (
	"web-server/database"
	"web-server/route"

	"github.com/gin-gonic/gin"
)

func main() {

	DB := database.Database(".env")
	defer DB.Close()
	database.Migertions(DB)

	router := gin.Default()
	route.RegisterRoute(router)
	router.Run(":8000")
}

//post user file or go route
//get by id
//get all
//delete
