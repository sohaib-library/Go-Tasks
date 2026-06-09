package repo

import "database/sql"

func InsertData(db *sql.DB, words, letters, spaces, lines, specialchar int) error {
	query := `INSERT INTO result (total_words, total_letters, total_spaces, total_lines, total_special_char) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, words, letters, spaces, lines, specialchar)
	if err != nil {
		return err
     }
	return err
}
