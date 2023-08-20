package model

// The user model
type User struct {
	// Generated user ID
	ID int `json:"id"`
	// User's first name / given name
	FirstName string `json:"firstname"`
	// User's last name / surname
	LastName string `json:"lastname"`
	// User's email address
	Email string `json:"email"`
}
