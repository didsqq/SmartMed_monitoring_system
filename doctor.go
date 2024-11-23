package smartmed

import "time"

type Doctor struct {
	Id             int       `json:"id" db:"id"`
	FullName       string    `json:"fullname" db:"fullname"`
	DateOfBirth    time.Time `json:"date_of_birth" db:"date_of_birth"`
	Email          string    `json:"email" db:"email"`
	Specialization string    `json:"specialization" db:"specialization"`
	PasswordHash   string    `json:"password" db:"password_hash"`
}
