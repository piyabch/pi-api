package model

// The authentication data model
type AuthData struct {
	// User's email address
	Email string `json:"email"`
	// User's password
	Password string `json:"password"`
}
