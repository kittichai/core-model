package dto

type UserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	UserName     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type UserResponse struct {
	ID        string          `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Email     string          `json:"email"`
	Message   MessageResponse `json:"message"`
}

type MessageResponse struct {
	TrackingID string `json:"tracking_id"`
	Message    string `json:"message"`
}
