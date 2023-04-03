package repository

import "sell-beauty-items/internal/users/model"

// users is a slice of User objects that will be used for authentication
var Users = []model.User{
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
}

type UserRepository struct{}

func (u *UserRepository) GetUsers() []model.User {
	return Users
}
