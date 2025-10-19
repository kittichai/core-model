package model

import "time"

type User struct {
	ID           string    `db:"id"`
	FirstName    string    `db:"firstname"`
	LastName     string    `db:"lastname"`
	UserName     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}
