package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

func CreateWallet(wallet request_response.Wallet) request_response.Wallet {
	wlt, err := models.CreateWallet(request_response.GetWalletModel(wallet))
	if err != nil {
		return request_response.Wallet{}
	}

	return request_response.GetWalletRequestResponse(wlt)
}

func GetWallet(userId string) request_response.Wallet {
	wlt, err := models.GetWallet(userId)
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

func UpdateWalletWithAmount(userId string, amount float64) error {
	wlt, err := models.GetWallet(userId)
	if err != nil { return err}

	wlt.Balance += amount
	wlt, err = models.UpdateWallet(wlt)
	return err
}

func UpdateWalletsForTransaction(txn models.Transaction) error {
	err := UpdateWalletWithAmount(txn.DestinationId, txn.Amount)
	if err != nil || txn.TransactionType == models.Load{ return err}

	err = UpdateWalletWithAmount(txn.SourceId, -txn.Amount)
	if err != nil { return err}
	return err
}