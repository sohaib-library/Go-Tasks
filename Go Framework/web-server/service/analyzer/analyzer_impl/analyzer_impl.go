package analyzer_service

import "database/sql"

type Analyzer_service struct {
	db *sql.DB
}

func NewAnalyzerImpl(db *sql.DB) *Analyzer_service {
	return &Analyzer_service{db: db}
}
