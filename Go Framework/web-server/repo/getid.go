package repo

import (
	"database/sql"
	"web-server/models"
)

// Get data by id
func GetResultByID(db *sql.DB, id int) (models.File, error) {
	var f models.File
	row := db.QueryRow(`
		SELECT id, total_words, total_letters, total_spaces, total_lines, total_special_char
		FROM result
		WHERE id = $1
	`, id)
	err := row.Scan(&f.ID,
			&f.TotalWords,
			&f.TotalLetters,
			&f.TotalSpaces,
			&f.TotalLines,
			&f.TotalSpecial)
	return f, err
}