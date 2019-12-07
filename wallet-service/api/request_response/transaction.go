package request_response

import "grabathon/models"

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

type Transaction struct {
	Id      int `json:"id"`
	SourceId  string `json:"source_id"`
	DestinationId  string `json:"destination_id"`
	TransactionType  string `json:"transaction_type"`
	Status string `json:"status"`
	Amount float64 `json:"amount"`
}

func GetTransactionRequestResponse(transaction models.Transaction) Transaction {
	txn := Transaction{
		Id: transaction.Id,
		SourceId:       transaction.SourceId,
		DestinationId:        transaction.DestinationId,
		Status:        transaction.Status,
		TransactionType: transaction.TransactionType,
		Amount: transaction.Amount,
	}

	return txn
}

func GetTransactionModel(transaction Transaction) models.Transaction {
	txn := models.Transaction{
		Id: transaction.Id,
		SourceId:       transaction.SourceId,
		DestinationId:        transaction.DestinationId,
		Status:        transaction.Status,
		TransactionType: transaction.TransactionType,
		Amount: transaction.Amount,
	}

	return txn
}

func GetTransactionRequestResponses(transactions []models.Transaction) []Transaction {
	var txns []Transaction
	for _, transaction := range transactions {
		txn := Transaction{
			Id:              transaction.Id,
			SourceId:        transaction.SourceId,
			DestinationId:   transaction.DestinationId,
			Status:          transaction.Status,
			TransactionType: transaction.TransactionType,
			Amount:          transaction.Amount,
		}

		txns = append(txns, txn)
	}
	return txns
}

func GetTransactionModels(transactions []Transaction) []models.Transaction {
	var txns []models.Transaction
	for _, transaction := range transactions {
		txn := models.Transaction{
			Id:              transaction.Id,
			SourceId:        transaction.SourceId,
			DestinationId:   transaction.DestinationId,
			Status:          transaction.Status,
			TransactionType: transaction.TransactionType,
			Amount:          transaction.Amount,
		}

		txns = append(txns, txn)
	}
	return txns
}
