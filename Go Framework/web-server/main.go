package main

import (
	"net/http"
    "web-server/Routines"

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

	router.GET("/count", filecount)

	router.Run(":8001")
}
func filecount(context *gin.Context) {
	letters, words, lines, spaces, special := routines.Filecontext()

	wordCountResponse:= WordCountResponse{
		WordCount: words,
		LetterCount: letters,
		LineCount: lines,
		SpaceCount: spaces,
		SpecialCharCount: special,
	}

	context.IndentedJSON(http.StatusOK, wordCountResponse)
}
