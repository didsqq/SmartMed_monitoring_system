package repository

import (
	"fmt"

	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/jmoiron/sqlx"
)

type AnalysisPostgres struct {
	db *sqlx.DB
}

func NewAnalysisPostgres(db *sqlx.DB) *AnalysisPostgres {
	return &AnalysisPostgres{db: db}
}

func (r *AnalysisPostgres) Create(userId int, analysis smartmed.Analysis) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (patient_id, pulse, respiratory_rate, oxygen_saturation, systolic_blood_pressure, diastolic_blood_pressure, heart_rate, analysis_timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", analysisTable)
	row := r.db.QueryRow(createQuery, analysis.PatientId, analysis.Pulse, analysis.RespiratoryRate, analysis.OxygenSaturation, analysis.SystolicBloodPressure, analysis.DiastolicBloodPressure, analysis.HeartRate, analysis.AnalysisTimestamp)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AnalysisPostgres) GetAll(patientId int) ([]smartmed.Analysis, error) {
	var analysis []smartmed.Analysis
	getAllQuery := fmt.Sprintf("SELECT * FROM %s WHERE patient_id=$1", analysisTable)
	err := r.db.Select(&analysis, getAllQuery, patientId)
	return analysis, err
}

func (r *AnalysisPostgres) GetById(analysisiId int) (smartmed.Analysis, error) {
	var analysis smartmed.Analysis
	getByIdQuery := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", analysisTable)
	err := r.db.Get(&analysis, getByIdQuery, analysisiId)
	return analysis, err
}
