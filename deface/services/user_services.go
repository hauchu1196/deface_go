package services

import (
	"deface/models"
	"errors"
)

var users []models.User
var idCounter uint = 1

func CreateUser(user models.User) models.User {
	user.ID = idCounter
	idCounter++
	users = append(users, user)
	return user
}

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id uint) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}
