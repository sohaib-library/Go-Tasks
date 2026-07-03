package analyzer_service

import (
	"web-server/models"
	"web-server/repo/analyzer/analyzer_impl"
)

func  (getbyid *Analyzer_service) GetByID( id, userID int) (models.File, error) {
	repo := analyzer_impl.NewAnalyzerImpl(getbyid.db)
	return repo.GetResultByID( id, userID)
}
