package repo

import (
	
	"web-server/database"
	"web-server/models"
)

func Login(login models.Login) (*models.Users, error) {
	// var user int64

	row := database.DB.QueryRow(`
		SELECT id , name , email , password
		FROM users 
		WHERE email = $1 
	`, login.EMAIL)

	var Resp models.Users

	err := row.Scan(
		&Resp.ID,
		&Resp.NAME,
		&Resp.EMAIL,
		&Resp.PASSWORD,
	)

	if err != nil {
		return nil, err
	}

	return &Resp, err
}
