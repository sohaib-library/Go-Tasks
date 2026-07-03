package analyzer_impl

import "database/sql"

type AnalyzerImpl struct {
	db *sql.DB
}

func NewAnalyzerImpl(db *sql.DB) *AnalyzerImpl {
	return &AnalyzerImpl{db: db}
}
