package service

import (
	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/repository"
)

type AnalysisService struct {
	repo repository.Analysis
}

func NewAnalysisService(repo repository.Analysis) *AnalysisService {
	return &AnalysisService{repo: repo}
}

func (s *AnalysisService) Create(userId int, analysis smartmed.Analysis) (int, error) {
	return s.repo.Create(userId, analysis)
}
