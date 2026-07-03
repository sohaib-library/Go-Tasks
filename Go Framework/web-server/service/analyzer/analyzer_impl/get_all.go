package analyzer_service

import (
	"web-server/models"
	"web-server/repo/analyzer/analyzer_impl"
)

func (getall *Analyzer_service) GetAll(userID int) ([]models.File, error) {

	repo := analyzer_impl.NewAnalyzerImpl(getall.db)

	return repo.GetAllResults(userID)

}
