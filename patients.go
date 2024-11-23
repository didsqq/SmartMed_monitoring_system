package smartmed

import "time"

type Patient struct {
	Id           int       `json:"id" db:"id"`
	DoctorId     int       `json:"doctor_id" db:"doctor_id"`
	FullName     string    `json:"fullname" db:"fullname"`
	DateOfBirth  time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender       string    `json:"gender" db:"gender"`
	Address      string    `json:"address" db:"address"`
	Email        string    `json:"email" db:"email"`
	PhoneNumber  string    `json:"phone_number" db:"phone_number"`
	PasswordHash string    `json:"password" db:"password_hash"`
}
