package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

type Transaction struct {
	Id      int
	SourceId  string
	DestinationId  string
	TransactionType  string
	Status string
	Amount float64
	PaymentId int
}

const(
	Pending = "PENDING"
	Completed = "COMPLETED"
	Invalid   = "INVALID"
)

const(
	TypePayment = "PAYMENT"
	RePayment = "REPAYMENT"
	Lend = "LEND"
	Load = "LOAD"
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
		if t.TransactionType == Lend{
			t.Status = Pending
		}
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


func GetTransactionsFromUser(uId string)([]Transaction, error) {
	var transactions []Transaction
	db.Where("source_id = ? and status = ?", uId, Pending).Find(&transactions)
	return transactions, nil
}

func GetTransactionsToUser(uId string)([]Transaction, error) {
	var transactions []Transaction
	db.Where("destination_id = ? and status = ?", uId, Pending).Find(&transactions)
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