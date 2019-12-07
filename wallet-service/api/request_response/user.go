package request_response

import (
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

type User struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	City        string `json:"city"`
	VouchCount  int `json:"vouch_count"`
	AutoPay     bool `json:"auto_pay"`
	Wallet      Wallet `json:"wallet"`
}

func GetUserRequestResponse(user models.User, w Wallet) User {
	usr := User{
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Name:        user.Name,
		City:        user.City,
		VouchCount:  user.VouchCount,
		AutoPay: user.AutoPay,
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
		AutoPay: user.AutoPay,
		VouchCount:  user.VouchCount,
	}

	return usr
}