package model

import "time"

type User struct {
	ID            string    `db:"id"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	UserName      string    `db:"username"`
	Email         string    `db:"email"`
	Phone         string    `db:"phone"`
	IsActive      bool      `db:"is_active"`
	EmailVerified bool      `db:"email_verified"`
	PasswordHash  string    `db:"password_hash"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
