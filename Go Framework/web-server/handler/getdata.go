package handler

import (
	
	"log"
	"web-server/repo"


	"github.com/gin-gonic/gin"
)

func Getdata(ctx *gin.Context) {

	row, err := repo.DB.Query(`

	SELECT id , total_words, total_letters, total_spaces, total_lines, total_special_char
	From result

	`)

	if err != nil {
		log.Panic("error" , err)
	}

	defer row.Close()

	






	
}
