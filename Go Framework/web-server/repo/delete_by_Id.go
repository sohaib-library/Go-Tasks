package repo

import "database/sql"

func DeleteResultByID(db *sql.DB, id, userID int) (int64, error) {
	res, err := db.Exec(`DELETE FROM result WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	return affected, err
}
