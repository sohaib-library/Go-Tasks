package analyzer

import (
	"web-server/models"
)

type AnalyzedServices interface {
	CreateById(userID, id int, filename string, content []byte) (*models.WordCountResponse, error)
	DeleteByID(id, userID int) (int64, error)
	GetAll(userID int) ([]models.File, error)
	GetByID(id, userID int) (models.File, error)
	UpdateByID(id, userID int, f models.File) (int64, error)
}