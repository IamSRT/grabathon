package request_response

import (
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type User struct {
	PhoneNumber string
	Email       string
	Name        string
	City        string
	VouchCount  int
	AutoPay     bool
	Wallet      Wallet
}

func GetUserRequestResponse(user models.User, w Wallet) User {
	usr := User{
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Name:        user.Name,
		City:        user.City,
		VouchCount:  user.VouchCount,
		Wallet:      w,
	}

	return usr
}

func GetUserModel(user User) models.User {
	usr := models.User{
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Name:        user.Name,
		City:        user.City,
		VouchCount:  user.VouchCount,
	}

	return usr
}