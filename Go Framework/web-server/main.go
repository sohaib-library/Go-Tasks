package main

import (
	"io"
	"net/http"
	"strconv"
	routines "web-server/Routines"

	"github.com/gin-gonic/gin"
)

type WordCountResponse struct {
	WordCount        int
	LetterCount      int
	LineCount        int
	SpaceCount       int
	SpecialCharCount int
}

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/count/:Id", filecount)

	router.Run(":8000")
}

func filecount(ctx *gin.Context) {



	key := ctx.Param("Id")
	Id, _ := strconv.Atoi(key)

	

	
	file , _, err := ctx.Request.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Failde upload file required "})

	}


	data , err := io.ReadAll(file)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}


	letters, words, lines, spaces, special := routines.Filecontext(Id, string(data))

	ctx.IndentedJSON(http.StatusOK, WordCountResponse{
		WordCount:        words,
		LetterCount:      letters,
		LineCount:        lines,
		SpaceCount:       spaces,
		SpecialCharCount: special,
	})
}
