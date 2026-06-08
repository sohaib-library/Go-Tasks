package repo

import (
	"database/sql"
	"web-server/models"
)

// Get all data
func GetAllResults(db *sql.DB) ([]models.File, error) {
	rows, err := db.Query(`
		SELECT id, total_words, total_letters, total_spaces, total_lines, total_special_char
		FROM result
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.File
	for rows.Next() {
		var f models.File
		if err := rows.Scan (
			&f.ID,
			&f.TotalWords,
			&f.TotalLetters,
			&f.TotalSpaces,
			&f.TotalLines,
			&f.TotalSpecial);
			 err != nil {
			return nil, err
		}
		results = append(results, f)
	}
	return results, rows.Err()
}