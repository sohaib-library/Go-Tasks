package auth_repo

import "database/sql"

type AuthImpl struct {
	db *sql.DB
}

func NewAuthImpl(db *sql.DB) *AuthImpl {
	return &AuthImpl{db: db}
}
