package handler

import (
	
	"web-server/service/analyzer"
	"web-server/service/auth"
)

type Handler struct {
	Authuser auth.AuthServices
	Analyzer analyzer.AnalyzedServices
}

func NewHandler(authuser auth.AuthServices, analyzer analyzer.AnalyzedServices) *Handler {
	return (&Handler{Authuser: authuser,
		Analyzer: analyzer})
}
