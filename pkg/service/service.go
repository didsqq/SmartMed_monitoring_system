package service

import (
	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/repository"
)

type Authorization interface {
}

type Analysis interface {
	Create(userId int, analysis smartmed.Analysis) (int, error)
}

type Service struct {
	Authorization
	Analysis
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Analysis: NewAnalysisService(repos.Analysis),
	}
}
