package repository

import (
	"fmt"

	models "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/jmoiron/sqlx"
)

type AnalysisPostgres struct {
	db *sqlx.DB
}

func NewAnalysisPostgres(db *sqlx.DB) *AnalysisPostgres {
	return &AnalysisPostgres{db: db}
}

func (r *AnalysisPostgres) Create(userId int, analysis models.Analysis) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (patient_id, pulse, respiratory_rate, oxygen_saturation, systolic_blood_pressure, diastolic_blood_pressure, heart_rate, analysis_timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", analysisTable)
	row := r.db.QueryRow(createQuery, analysis.PatientID, analysis.Pulse, analysis.RespiratoryRate, analysis.OxygenSaturation, analysis.SystolicBloodPressure, analysis.DiastolicBloodPressure, analysis.HeartRate, analysis.AnalysisTimestamp)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
