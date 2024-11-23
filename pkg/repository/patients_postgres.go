package repository

import (
	"fmt"

	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/jmoiron/sqlx"
)

type PatientsPostgres struct {
	db *sqlx.DB
}

func NewPatientsPostgres(db *sqlx.DB) *PatientsPostgres {
	return &PatientsPostgres{db: db}
}

func (r *PatientsPostgres) Create(userId int, patinet smartmed.Patient) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (doctor_id, fullname, date_of_birth, gender, address, email, phone_number, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", patientsTable)
	row := r.db.QueryRow(createQuery, patinet.DoctorId, patinet.FullName, patinet.DateOfBirth, patinet.Gender, patinet.Address, patinet.Email, patinet.PhoneNumber, patinet.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PatientsPostgres) GetAll(userId int) ([]smartmed.Patient, error) {
	var patients []smartmed.Patient
	getListsQuery := fmt.Sprintf("SELECT * FROM %s WHERE doctor_id=$1", patientsTable)
	err := r.db.Select(&patients, getListsQuery, userId)
	return patients, err
}
