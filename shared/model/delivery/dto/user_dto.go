package dto

import "time"

type UserRequest struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"password_hash"`
}

type UserResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       Data   `json:"data"`
	Error      Error  `json:"error"`
	TrackingID string `json:"tracking_id"`
}

type UserErrorResponse struct {
	Status     string `json:"status"`
	Error      Error  `json:"error"`
	TrackingID string `json:"tracking_id"`
}

type Data struct {
	ID            string    `json:"id"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	UserName      string    `json:"username"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	IsActive      bool      `json:"is_active"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	//Detail  Detail `json:"detail"`
}

type Detail struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"password_hash"`
}

/*
{
    "code": "VALIDATION_FAILED",
    "message": "One or more validation errors occurred.",
    "details": {
      "email": "Email already exists.",
      "password": "Password must be stronger."
    }


{
  "status": "success",
  "message": "User registered successfully.",
  "data": {
    "id": "e0a89f13-3b7c-4a1c-a2a1-6e5d1b3b4a2a",
    "email": "user@example.com",
    "username": "john_doe",
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+66912345678",
    "created_at": "2025-10-20T08:31:00Z"
  }
}

{
  "status": "error",
  "error": {
    "code": "VALIDATION_FAILED",
    "message": "One or more validation errors occurred.",
    "details": {
      "email": "Email already exists.",
      "password": "Password must be stronger."
    }
  },
  "trace_id": "a7d9f1f8-5c6b-4a2b-9d31-bbd4f3f1c3a1"
}

*/
