package service

import (
	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/repository"
)

type PatientsService struct {
	repo repository.Patients
}

func NewPatientsService(repo repository.Patients) *PatientsService {
	return &PatientsService{repo: repo}
}

func (s *PatientsService) GetAll(userId int) ([]smartmed.Patient, error) {
	return s.repo.GetAll(userId)
}

func (s *PatientsService) Create(userId int, patinet smartmed.Patient) (int, error) {
	// patinet.PasswordHash = s.generatePasswordHash(patinet.PasswordHash)
	return s.repo.Create(userId, patinet)
}
