package request_response

import "grabathon/models"

/**
 * Created by Sai Ravi Teja K on 07, Dec 2019
 * © Bundl Technologies Private Ltd.
 */

type Payment struct {
	Id int
	Status string
	Type string
	Transactions []Transaction
}

func GetPaymentRequestResponse(payment models.Payment, transactions []models.Transaction) Payment {
	pmt := Payment{
		Id:           payment.Id,
		Status:       payment.Status,
		Type:         payment.Type,
		Transactions: nil,
	}

	pmt.Transactions = GetTransactionRequestResponses(transactions)

	return pmt
}

func GetPaymentModel(payment Payment) (models.Payment, []models.Transaction) {
	pmt := models.Payment{
		Id:           payment.Id,
		Status:       payment.Status,
		Type:         payment.Type,
	}

	return pmt,GetTransactionModels(payment.Transactions)
}