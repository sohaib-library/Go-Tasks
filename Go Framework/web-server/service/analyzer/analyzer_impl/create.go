package analyzer_service

import (
	"fmt"
	"log"
	"strings"
	"web-server/models"
	"web-server/repo/analyzer/analyzer_impl"
)

func (a *Analyzer_service) CreateById(userID, id int, filename string, content []byte) (*models.WordCountResponse, error) {

	if !strings.HasSuffix(filename, ".txt") {
		return nil, fmt.Errorf("only .txt files allowed")
	}

	letters, words, lines, spaces, special := Filecontext(id, string(content))

	repo := analyzer_impl.NewAnalyzerImpl(a.db)
	err := repo.Create(userID, words, letters, spaces, lines, special)
	if err != nil {
		log.Print("error : Internal Server Error")
		return nil, err
	}

	return &models.WordCountResponse{
		WordCount:        words,
		LetterCount:      letters,
		LineCount:        lines,
		SpaceCount:       spaces,
		SpecialCharCount: special,
	}, nil
}
