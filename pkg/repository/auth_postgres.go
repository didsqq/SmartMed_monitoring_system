package repository

import (
	"fmt"

	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateDoctor(doctor smartmed.Doctor) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (fullname, date_of_birth, email, specialization, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id", doctorsTable)
	row := r.db.QueryRow(createQuery, doctor.FullName, doctor.DateOfBirth, doctor.Email, doctor.Specialization, doctor.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetDoctor(email, password string) (smartmed.Doctor, error) {
	var doctor smartmed.Doctor
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", doctorsTable)
	err := r.db.Get(&doctor, getQuery, email, password)
	return doctor, err
}
