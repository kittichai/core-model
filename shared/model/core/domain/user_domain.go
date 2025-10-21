package domain

import "time"

type User struct {
	ID            string
	FirstName     string
	LastName      string
	UserName      string
	Email         string
	Phone         string
	IsActive      bool
	EmailVerified bool
	PasswordHash  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
