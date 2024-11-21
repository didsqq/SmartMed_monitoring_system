package repository

import (
	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Analysis interface {
	Create(userId int, analysis smartmed.Analysis) (int, error)
}

type Repository struct {
	Analysis
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Analysis: NewAnalysisPostgres(db),
	}
}
