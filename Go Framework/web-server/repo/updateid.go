package repo

import (
	"database/sql"
	"web-server/models"
)


func UpdateResultByID(db *sql.DB, id int, f models.File) (int64, error) {
	res, err := db.Exec(`
		UPDATE result
		SET total_words        = $1,
		    total_letters      = $2,
		    total_spaces       = $3,
		    total_lines        = $4,
		    total_special_char = $5
		WHERE id = $6
	`, f.TotalWords, f.TotalLetters, f.TotalSpaces, f.TotalLines, f.TotalSpecial, id)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	return affected, err
}
