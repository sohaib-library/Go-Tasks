package analyzer_service

import (
	"web-server/repo/analyzer/analyzer_impl"
)

func (delete *Analyzer_service) DeleteByID(id, userID int) (int64, error) {

	repo := analyzer_impl.NewAnalyzerImpl(delete.db)

	return repo.DeleteResultByID(id, userID)

}
