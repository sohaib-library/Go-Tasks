package service

import (
	"database/sql"
	"log"
	"web-server/models"
	"web-server/repo"

	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB, login models.Login )  {

	Resp ,err := repo.Login(login)
	if err != nil {
		log.Print(err)
		
	}

	userPass:= []byte(Resp.PASSWORD)
    dbPass:= []byte(login.PASSWORD)
    passErr:= bcrypt.CompareHashAndPassword(dbPass, userPass)
    if passErr != nil{
    log.Println(passErr)    
       
    }










}