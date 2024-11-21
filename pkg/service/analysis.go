package service

import (
	models "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/repository"
)

type AnalysisService struct {
	repo repository.Analysis
}

func NewAnalysisService(repo repository.Analysis) *AnalysisService {
	return &AnalysisService{repo: repo}
}

func (s *AnalysisService) Create(userId int, analysis models.Analysis) (int, error) {
	return s.repo.Create(userId, analysis)
}
