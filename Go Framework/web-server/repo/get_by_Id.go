package repo

import (
	"database/sql"
	"web-server/models"
)

// Get data by id
func GetResultByID(db *sql.DB, id, userID int) (models.File, error) {
	var f models.File
	row := db.QueryRow(`
		SELECT id, user_id, total_words, total_letters, total_spaces, total_lines, total_special_char
		FROM result
		WHERE id = $1 AND user_id = $2
	`, id, userID)
	err := row.Scan(&f.ID,
		&f.UserID,
		&f.TotalWords,
		&f.TotalLetters,
		&f.TotalSpaces,
		&f.TotalLines,
		&f.TotalSpecial)
	return f, err
}
