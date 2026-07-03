package analyzer

import (
	"web-server/models"

	"github.com/gin-gonic/gin"
)

type AnalyzedServices interface {
	CreateById(ctx *gin.Context) (int, any)
	DeleteByID(id, userID int) (int64, error)
	GetAll(userID int) ([]models.File, error)
	GetByID(id, userID int) (models.File, error)
	UpdateByID(id, userID int, f models.File) (int64, error)
}