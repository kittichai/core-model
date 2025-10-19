package dto

type UserRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserResponse struct {
	ID      string          `json:"id"`
	Email   string          `json:"email"`
	Name    string          `json:"name"`
	Message MessageResponse `json:"message"`
}

type MessageResponse struct {
	TrackingID string `json:"tracking_id"`
	Message    string `json:"message"`
}
