package service

import (
	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/repository"
)

type Authorization interface {
	CreateDoctor(smartmed.Doctor) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accesToken string) (int, error)
}

type Analysis interface {
	Create(userId int, analysis smartmed.Analysis) (int, error)
	GetAll() ([]smartmed.Analysis, error)
	GetById(analysisiId int) (smartmed.Analysis, error)
}

type Patients interface {
	Create(userId int, patinet smartmed.Patient) (int, error)
	GetAll(userId int) ([]smartmed.Patient, error)
}

type Service struct {
	Authorization
	Analysis
	Patients
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Analysis:      NewAnalysisService(repos.Analysis),
		Authorization: NewAuthService(repos.Authorization),
		Patients:      NewPatientsService(repos.Patients),
	}
}
