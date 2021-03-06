package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
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
	v := GetAllVouchesForVouchee(usr.PhoneNumber)
	usr.VouchCount = len(v)
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

func EnableAutoPay(vouches []request_response.Vouch) []request_response.Vouch {
	for i := range vouches{
		vouches[i].VouchType = models.AutoPay
	}
	usr, err := models.GetUser(vouches[0].VoucherId)
	if err != nil { return nil }

	usr.AutoPay = true
	usr, err =models.UpdateUser(usr)
	if err != nil { return nil }
	return CreateVouches(vouches)
}

func Vouch(vouches []request_response.Vouch) []request_response.Vouch {
	for i := range vouches{
		vouches[i].VouchType = models.Default
	}
	return CreateVouches(vouches)
}

func IsAutoEnabled(voucheeId string, voucherId string) bool {
	f := GetUser(voucherId).AutoPay
	f = f && models.IsVouch(voucheeId, voucherId, models.AutoPay)
	return f
}

func IsVouchValid(voucheeId string, voucherId string, amount float64) bool {
	wlt, err := models.GetWallet(voucherId)
	if err != nil {
		return false
	}

	if wlt.Balance > amount {
		return true
	}
	return false
}