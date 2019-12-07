package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

func CreateUser(user request_response.User) request_response.User {
	usr, err := models.CreateUser(request_response.GetUserModel(user))
	if err != nil {
		return request_response.User{}
	}
	w := GetWallet(usr.PhoneNumber)
	return request_response.GetUserRequestResponse(usr, w)
}

func GetAllUsers() []request_response.User {
	_, err := models.GetAllUsers()
	if err != nil {
		return []request_response.User{}
	}
	return []request_response.User{}
}

func GetUser(id string) request_response.User {
	usr, err := models.GetUser(id)
	if err != nil {
		return request_response.User{}
	}
	w := GetWallet(usr.PhoneNumber)
	return request_response.GetUserRequestResponse(usr,w)
}

func UpdateUser(user request_response.User) request_response.User {
	usr, err := models.UpdateUser(request_response.GetUserModel(user))
	if err != nil {
		return request_response.User{}
	}
	w := GetWallet(usr.PhoneNumber)
	return request_response.GetUserRequestResponse(usr,w)
}