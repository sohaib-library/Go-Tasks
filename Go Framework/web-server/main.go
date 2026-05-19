package main

import (
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

	letters, words, lines, spaces, special := routines.Filecontext(Id)

	wordCountResponse := WordCountResponse{
		WordCount:        words,
		LetterCount:      letters,
		LineCount:        lines,
		SpaceCount:       spaces,
		SpecialCharCount: special,
	}

	ctx.IndentedJSON(http.StatusOK, wordCountResponse)
}
