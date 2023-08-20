package model

// The authentication result model
type AuthResult struct {
	// User's email address
	Token string `json:"token"`
}
