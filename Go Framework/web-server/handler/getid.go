package handler

import (
	// "log"
	// "net/http"
	"strconv" 
	"web-server/models"
	"web-server/repo"

	"github.com/gin-gonic/gin"
)

func Getid(ctx *gin.Context) {

	// Get Specific Id from user 

	key := ctx.Param("Id")
	Id, _ := strconv.Atoi(key)
	


    row := repo.DB.QueryRow(`

	SELECT id , total_words, total_letters, total_spaces, total_lines, total_special_char
	From result

	WHERE id = $1

	`, Id )

	
		var f models.File
		err := row.Scan(

            &f.ID,
            &f.TotalWords,
            &f.TotalLetters,
            &f.TotalSpaces,
            &f.TotalLines,
            &f.TotalSpecial,
		)


		if err != nil {

			ctx.JSON (500 , gin.H{"error" : err.Error()})

		}
	



	ctx.JSON(200 , f)


}
