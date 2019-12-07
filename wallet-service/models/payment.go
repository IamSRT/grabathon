package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

type Payment struct {
	Id int
	Status string
	Type string
}

func CreatePayment(p Payment) (Payment, error) {
	if err := db.Create(&p).Error; err != nil {
		return Payment{}, err
	}
	return p, nil
}

func CreatePayments(payments []Payment) ([]Payment, error) {
	var ps []Payment
	for _, p := range payments {
		db.Create(&p)
		ps = append(ps, p)
	}
	return ps, nil
}

func GetPayment(pId int)(Payment, error) {
	var payment Payment
	db.Find(&payment, pId)
	return payment, nil
}

func UpdatePayment(payment Payment)(Payment, error){
	if err := db.Model(&payment).Updates(payment).Error; err != nil {
		return Payment{}, err
	}
	return payment, nil
}

func DeletePayment(payment Payment) error {
	db.Delete(&payment)
	return nil
}