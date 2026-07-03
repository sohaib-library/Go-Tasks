package analyzer_service

import (
	"web-server/models"
	"web-server/repo/analyzer/analyzer_impl"
)

func (update *Analyzer_service) UpdateByID(id, userID int, f models.File) (int64, error) {

	repo := analyzer_impl.NewAnalyzerImpl(update.db)

	return repo.UpdateResultByID(id, userID, f)
}
