package repo

import (
	"database/sql"
	"web-server/models"
)

// Update id
func UpdateResultByID(db *sql.DB, id, userID int, f models.File) (int64, error) {
	res, err := db.Exec(`
		UPDATE result
		SET total_words        = $1,
		    total_letters      = $2,
		    total_spaces       = $3,
		    total_lines        = $4,
		    total_special_char = $5
		WHERE id = $6 AND user_id = $7
	`, f.TotalWords,
		f.TotalLetters,
		f.TotalSpaces,
		f.TotalLines,
		f.TotalSpecial,
		id,
		userID)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, err

}
