package repository

import (
	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateDoctor(doctor smartmed.Doctor) (int, error)
	GetDoctor(email, password string) (smartmed.Doctor, error)
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

type Repository struct {
	Analysis
	Authorization
	Patients
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Analysis:      NewAnalysisPostgres(db),
		Authorization: NewAuthPostgres(db),
		Patients:      NewPatientsPostgres(db),
	}
}
