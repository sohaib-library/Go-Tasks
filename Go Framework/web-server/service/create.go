package service

import (
	"database/sql"
	"web-server/models"
	"web-server/repo"
)


func CreateWordCount(db *sql.DB, workers int, content string) (models.WordCountResponse, error) {
	letters, words, lines, spaces, special := Filecontext(workers, content)

	err := repo.InsertResult(db, words, letters, spaces, lines, special)
	if err != nil {
		return models.WordCountResponse{}, err
	}

	return models.WordCountResponse{
		WordCount:        words,
		LetterCount:      letters,
		LineCount:        lines,
		SpaceCount:       spaces,
		SpecialCharCount: special,
	}, nil
}
