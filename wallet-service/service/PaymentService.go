package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 07, Dec 2019
 * © Bundl Technologies Private Ltd.
 */

func CreatePayment(payment request_response.Payment) request_response.Payment {
	p, t := request_response.GetPaymentModel(payment)
	p, err := models.CreatePayment(p)
	if err != nil {
		return request_response.Payment{}
	}

	txns, err := models.CreateTransactions(t)
	if err != nil {
		return request_response.Payment{}
	}
	for _,t := range txns {
		err = UpdateWalletsForTransaction(t)
		if err != nil {
			return request_response.Payment{}
		}
	}
	return request_response.GetPaymentRequestResponse(p, txns)
}

func GetPayment(pId int) request_response.Payment {
	p, err := models.GetPayment(pId)
	if err != nil {
		return request_response.Payment{}
	}

	txns, err := models.GetTransactions(p.Id)
	return request_response.GetPaymentRequestResponse(p, txns)
}

func DeletePayment(pId int) error {
	err := models.DeletePayment(models.Payment{
		Id:     pId,
	})
	if err != nil {
		return err
	}

	txns, err := models.GetTransactions(pId)
	for _,t := range txns{
		err = models.DeleteTransaction(t)
		if err != nil {
			return err
		}
	}
	return nil
}