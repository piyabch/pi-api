package db

import (
	"database/sql"
	"fmt"

	"github.com/piyabch/pi-api/model"
)

// Insert a user to a DB.
// Return the auto generated user ID from the DB.
func CreateUser(user model.User) (int, error) {
	result, err := db.Exec("INSERT INTO user (firstname, lastname, email) VALUES (?, ?, ?)",
		user.FirstName, user.LastName, user.Email)
	if err != nil {
		return 0, fmt.Errorf("SQL Error: %v", err)
	}
	id, err := result.LastInsertId()
	if id <= 0 || err != nil {
		return 0, fmt.Errorf("SQL Error: %v", err)
	}
	return int(id), nil
}

// Find a user by the unique user ID.
// Return the matching user.
func FindUserByID(id int) (model.User, error) {
	var user model.User

	row := db.QueryRow("SELECT id, firstname, lastname, email FROM user WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("Data not found, User ID: %d", id)
		}
		return user, fmt.Errorf("SQL Error: %v, User ID: %d", err, id)
	}
	return user, nil
}

// Find a user which the firstname
// contains a input name string.
// Return the matching users.
func FindUsersByName(name string) ([]model.User, error) {
	var users []model.User
	searchName := "%" + name + "%"
	rows, err := db.Query("SELECT id, firstname, lastname, email FROM user WHERE firstname like ?", searchName)
	if err != nil {
		return nil, fmt.Errorf("SQL Error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, fmt.Errorf("SQL Error: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SQL Error: %v", err)
	}
	return users, nil
}

// Update user by the user ID.
// The updatable fields are as follows.
//   - Firstname
//   - Lastname
//   - Email
//
// Return the updated row count.
func UpdateUser(user model.User) (int, error) {
	result, err := db.Exec("UPDATE user SET firstname=?, lastname=?, email=? WHERE id=?",
		user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return 0, fmt.Errorf("SQL Error: %v", err)
	}
	rowCount, err := result.RowsAffected()
	if rowCount <= 0 || err != nil {
		return 0, fmt.Errorf("Update failed - User not found, User ID: %d", user.ID)
	}
	return user.ID, nil
}
