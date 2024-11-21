package smartmed

import "time"

type Analysis struct {
	Id                     int       `json:"id" db:"id"`
	PatientId              int       `json:"patient_id" db:"patient_id"`
	Pulse                  int16     `json:"pulse" db:"pulse"`
	RespiratoryRate        int16     `json:"respiratory_rate" db:"respiratory_rate"`
	OxygenSaturation       float64   `json:"oxygen_saturation" db:"oxygen_saturation"`
	SystolicBloodPressure  int       `json:"systolic_blood_pressure" db:"systolic_blood_pressure"`
	DiastolicBloodPressure int       `json:"diastolic_blood_pressure" db:"diastolic_blood_pressure"`
	HeartRate              int       `json:"heart_rate" db:"heart_rate"`
	AnalysisTimestamp      time.Time `json:"analysis_timestamp" db:"analysis_timestamp"`
}
