package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

func CreateWallet(wallet request_response.Wallet) request_response.Wallet {
	wlt, err := models.CreateWallet(request_response.GetWalletModel(wallet))
	if err != nil {
		return request_response.Wallet{}
	}

	return request_response.GetWalletRequestResponse(wlt)
}

func GetWallet(id int) request_response.Wallet {
	wlt, err := models.GetWallet(id)
	if err != nil {
		return request_response.Wallet{}
	}
	return request_response.GetWalletRequestResponse(wlt)
}

func UpdateWallet(wallet request_response.Wallet) request_response.Wallet {
	wlt, err := models.UpdateWallet(request_response.GetWalletModel(wallet))
	if err != nil {
		return request_response.Wallet{}
	}

	return request_response.GetWalletRequestResponse(wlt)
}