package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

func CreateTransaction(wallet request_response.Transaction) request_response.Transaction {
	txn, err := models.CreateTransaction(request_response.GetTransactionModel(wallet))
	if err != nil {
		return request_response.Transaction{}
	}

	return request_response.GetTransactionRequestResponse(txn)
}

func GetTransaction(txnId int) request_response.Transaction {
	txn, err := models.GetTransaction(txnId)
	if err != nil {
		return request_response.Transaction{}
	}
	return request_response.GetTransactionRequestResponse(txn)
}

func UpdateTransaction(tran request_response.Transaction) request_response.Transaction {
	txn, err := models.UpdateTransaction(request_response.GetTransactionModel(tran))
	if err != nil {
		return request_response.Transaction{}
	}

	return request_response.GetTransactionRequestResponse(txn)
}

func DeleteTransaction(txnId int) error {
	err := models.DeleteTransaction(models.Transaction{
		Id: txnId,
	})
	return err
}