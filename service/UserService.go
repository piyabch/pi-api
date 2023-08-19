package service

import (
	"github.com/piyabch/pi-api/db"
	"github.com/piyabch/pi-api/model"
)

func CreateUser(user *model.User) (*model.User, error) {
	id, err := db.CreateUser(*user)
	if id > 0 && err == nil {
		user.ID = int(id)
	}
	return user, nil
}

func FindUserByID(id int) (*model.User, error) {
	user, err := db.FindUserByID(id)
	return &user, err
}

func FindUsersByName(name string) ([]model.User, error) {
	users, err := db.FindUsersByName(name)
	return users, err
}

func UpdateUser(user *model.User) (*model.User, error) {
	db.UpdateUser(*user)
	return user, nil
}
