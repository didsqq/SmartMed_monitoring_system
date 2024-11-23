package repository

import (
	"fmt"

	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PatientsPostgres struct {
	db *sqlx.DB
}

func NewPatientsPostgres(db *sqlx.DB) *PatientsPostgres {
	return &PatientsPostgres{db: db}
}

func (r *PatientsPostgres) Create(userId int, patinet smartmed.Patient) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (doctor_id, fullname, date_of_birth, gender, address, email, phone_number,chatid, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", patientsTable)
	row := r.db.QueryRow(createQuery, userId, patinet.FullName, patinet.DateOfBirth, patinet.Gender, patinet.Address, patinet.Email, patinet.PhoneNumber, "", patinet.PasswordHash)
	logrus.Info(userId)
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

func (r *PatientsPostgres) GetById(userId int) (smartmed.Patient, error) {
	var patient smartmed.Patient
	getByIdQuery := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", patientsTable)
	err := r.db.Get(&patient, getByIdQuery, userId)
	return patient, err
}

func (r *PatientsPostgres) GetChatId(patientId int) (string, error) {
	var patient smartmed.Patient
	getByIdQuery := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", patientsTable)
	err := r.db.Get(&patient, getByIdQuery, patientId)

	return patient.ChatId, err
	// var chatId string
	// query := fmt.Sprintf("SELECT chatid FROM %s WHERE id=$1", patientsTable)

	// // Используем метод Get для получения одного значения
	// err := r.db.Get(&chatId, query, patientId)
	// if err != nil {
	// 	return -1, err // Возвращаем пустую строку и ошибку
	// }

	// res, err := strconv.Atoi(chatId)
	// if err != nil {
	// 	return -1, err // Возвращаем пустую строку и ошибку
	// }
	// return res, nil
}
