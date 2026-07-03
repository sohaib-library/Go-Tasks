package analyzer_impl

import (
	"web-server/models"
)

// Get all data
func (data *AnalyzerImpl) GetAllResults(userID int) ([]models.File, error) {

	rows, err := data.db.Query(`
		SELECT id,
		user_id, 
		total_words,
		total_letters,
	    total_spaces, 
		total_lines,
		 total_special_char
		FROM result
		WHERE user_id = $1
		ORDER BY id
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []models.File

	for rows.Next() {
		var f models.File
		if err := rows.Scan(
			&f.ID,
			&f.UserID,
			&f.TotalWords,
			&f.TotalLetters,
			&f.TotalSpaces,
			&f.TotalLines,
			&f.TotalSpecial); err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	return results, rows.Err()
}
