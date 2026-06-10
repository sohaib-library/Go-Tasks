package repo

import "database/sql"

func SignUP(db *sql.DB ,id int , name string , email string , password string ) error {
	query := `INSERT INTO users (id , name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query , id , name , email ,password)
	if err != nil {
		return err
     }
	return err
}
