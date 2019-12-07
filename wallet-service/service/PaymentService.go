package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 07, Dec 2019
 * © Refugee Inc
 */

func CreatePayment(payment request_response.Payment) request_response.Payment {
	p, txns := request_response.GetPaymentModel(payment)
	p, err := models.CreatePayment(p)
	if err != nil {
		return request_response.Payment{}
	}

	for i := range txns{
		txns[i].PaymentId = p.Id
	}
	txns, err = models.CreateTransactions(txns)
	if err != nil {
		return request_response.Payment{}
	}

	var amount float64
	for _,t := range txns {
		err = UpdateWalletsForTransaction(t)
		if err != nil {
			return request_response.Payment{}
		}

		amount += t.Amount
	}

	if payment.Type == models.RePayment{
		Settle(payment, amount)
	}

	return request_response.GetPaymentRequestResponse(p, txns)
}

func Settle(payment request_response.Payment, amount float64) {
	if len(payment.Transactions) != 1 {
		return
	}
	balance := GetBalances(payment.Transactions[0].DestinationId)
	if _, ok := balance[payment.Transactions[0].SourceId]; !ok {
		return
	}

	txns, _ := models.GetTransactionsFromUser(payment.Transactions[0].DestinationId)
	for _, t := range txns {
		if amount < t.Amount {
			continue
		}
		amount -= t.Amount
		t.Status = models.Completed
		_, _ = models.UpdateTransaction(t)
	}

	if amount > 0 {
		for _, t := range txns {
			if t.Status == models.Completed {
				continue
			}
			t.Amount = t.Amount - amount
			_, _ = models.UpdateTransaction(t)
		}
	}
}

func GetPayment(pId int) request_response.Payment {
	p, err := models.GetPayment(pId)
	if err != nil {
		return request_response.Payment{}
	}

	txns, err := models.GetTransactions(p.Id)
	return request_response.GetPaymentRequestResponse(p, txns)
}

func GetPendingPayments(userId string) request_response.CompleteWallet {
	cw := request_response.CompleteWallet{
		WalletId: 0,
		UserId:   userId,
		Amount:   0,
		Balances: nil,
	}

	w := GetWallet(userId)
	cw.Amount = w.Balance
	cw.WalletId = w.Id

	balances := GetBalances(userId)
	var b []request_response.Balance
	for id, amt := range balances {
		b = append(b, request_response.Balance{
			UserId:  id,
			Balance: amt,
		})
	}

	cw.Balances = b

	return cw
}

func GetBalances(userId string) map[string]float64 {
	balances := make(map[string]float64)
	txns, _ := models.GetTransactionsFromUser(userId)
	for _, txn := range txns {
		if _, ok := balances[txn.DestinationId]; !ok {
			balances[txn.DestinationId] = 0
		}
		balances[txn.DestinationId] += txn.Amount
	}
	txns, _ = models.GetTransactionsToUser(userId)
	for _, txn := range txns {
		if _, ok := balances[txn.SourceId]; !ok {
			balances[txn.SourceId] = 0
		}
		balances[txn.SourceId] -= txn.Amount
	}
	return balances
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