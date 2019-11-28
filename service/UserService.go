package service

import (
	"grabathon/api/request"
	. "grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

func CreateUser(user request.User) User {
	usr, err := Create(User{})
	if err != nil {
		return User{}
	}
	return usr
}

func GetAllUsers() []User {
	return nil
}

func GetUser() User {
	return User{}
}