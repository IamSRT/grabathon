package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type Transaction struct {
	Id      int
	SourceId  string
	DestinationId  string
	TransactionType  string
	Status string
	Amount float64
}

const(
	Initiated = "INITIATED"
	Completed = "COMPLETED"
	Invalid   = "INVALID"
)

const(
	TypePayment = "PAYMENT"
	Lend = "LEND"
)

func CreateTransaction(t Transaction) (Transaction, error) {
	if err := db.Create(&t).Error; err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func CreateTransactions(transactions []Transaction) ([]Transaction, error) {
	var txns []Transaction
	for _, t := range transactions {
		db.Create(&t)
		txns = append(txns, t)
	}
	return txns, nil
}

func GetTransaction(tId int)(Transaction, error) {
	var transaction Transaction
	db.Find(&transaction, tId)
	return transaction, nil
}

func GetTransactions(pId int)([]Transaction, error) {
	var transactions []Transaction
	db.Where("payment_id = ?", pId).Find(&transactions)
	return transactions, nil
}

func UpdateTransaction(transaction Transaction)(Transaction, error){
	if err := db.Model(&transaction).Updates(transaction).Error; err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func DeleteTransaction(transaction Transaction) error {
	db.Delete(&transaction)
	return nil
}