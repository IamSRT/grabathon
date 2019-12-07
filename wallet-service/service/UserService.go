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

func EnableAutoPay(userId string, autoPayees []string) []request_response.Vouch {
	var vouches []request_response.Vouch
	for _,autoPayee := range autoPayees{
		vouches = append(vouches, request_response.Vouch{
			VoucheeId: autoPayee,
			VoucherId: userId,
			VouchType: models.AutoPay,
		})
	}

	return CreateVouches(vouches)
}

func Vouch(userId string, autoPayees []string) []request_response.Vouch {
	var vouches []request_response.Vouch
	for _,autoPayee := range autoPayees{
		vouches = append(vouches, request_response.Vouch{
			VoucheeId: autoPayee,
			VoucherId: userId,
			VouchType: models.Default,
		})
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