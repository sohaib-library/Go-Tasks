package analyzer_service

import (
	"errors"
	"web-server/repo/analyzer/analyzer_impl"

	"github.com/sirupsen/logrus"
)

func (delete *Analyzer_service) DeleteByID(id, userID int) (int64, error) {

	repo := analyzer_impl.NewAnalyzerImpl(delete.db)

	affected, err := repo.DeleteResultByID(id, userID)
	if err != nil {
		logrus.Error(err)
		return 0, errors.New("failed to delete user")
	}

	if affected == 0 {
		return 0, errors.New("user not found")
	}

	return affected, nil
}
