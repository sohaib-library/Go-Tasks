package repo

import "database/sql"

func SignUP(db *sql.DB , name string , email string , password string ) (int64, error) {
	query := `INSERT INTO users ( name, email, password) VALUES ($1, $2, $3)`
	res , err := db.Exec(query , name , email , password)
	if err != nil {
		return 0, err
     }
	affected , err := res.RowsAffected()
	return affected , err
}
