package repo

import "database/sql"


func DeleteResultByID(db *sql.DB, id int) (int64, error) {
	res, err := db.Exec(`DELETE FROM result WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	return affected, err
}
