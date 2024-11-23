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

func (s *AnalysisService) GetAll(patientId int) ([]smartmed.Analysis, error) {
	return s.repo.GetAll(patientId)
}

func (s *AnalysisService) GetById(analysisiId int) (smartmed.Analysis, error) {
	return s.repo.GetById(analysisiId)
}
