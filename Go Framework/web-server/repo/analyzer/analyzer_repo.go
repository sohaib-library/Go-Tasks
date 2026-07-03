package analyze_repo

import (
	"database/sql"
	"web-server/models"
)


type AnalyzerRepo interface {
	Create(db *sql.DB, userID, words, letters, spaces, lines, specialchar int) error
	DeleteByID(db *sql.DB, id, userID int) (int64, error)
	GetAllData(db *sql.DB, userID int) ([]models.File, error)
	GetByID(db *sql.DB, id, userID int) (models.File, error)
	UpdateByID(db *sql.DB, id, userID int, f models.File) (int64, error)
}