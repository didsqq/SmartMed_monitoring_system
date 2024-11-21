package smartmed

import "time"

type Analysis struct {
	Id                     int       `json:"id"`
	PatientId              int       `json:"patient_id"`
	Pulse                  int16     `json:"pulse"`
	RespiratoryRate        int16     `json:"respiratory_rate"`
	OxygenSaturation       float64   `json:"oxygen_saturation"`
	SystolicBloodPressure  int       `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int       `json:"diastolic_blood_pressure"`
	HeartRate              int       `json:"heart_rate"`
	AnalysisTimestamp      time.Time `json:"analysis_timestamp"`
}
