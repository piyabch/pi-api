package service

import (
	"github.com/piyabch/pi-api/db"
	"github.com/piyabch/pi-api/model"
)

// Create a new user.
// Fill the given user ID on success creation.
// Return the newly created user.
func CreateUser(user *model.User) (*model.User, error) {
	id, err := db.CreateUser(*user)
	if id > 0 && err == nil {
		user.ID = int(id)
	}
	return user, nil
}

// Find user by ID.
// Return the user found.
func FindUserByID(id int) (*model.User, error) {
	user, err := db.FindUserByID(id)
	return &user, err
}

// Find users by Name
// Return the array of users found.
func FindUsersByName(name string) ([]model.User, error) {
	users, err := db.FindUsersByName(name)
	return users, err
}

// Update the user data
// Return the updated user.
func UpdateUser(user *model.User) (*model.User, error) {
	db.UpdateUser(*user)
	return user, nil
}
