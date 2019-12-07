package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type Wallet struct {
	Id      int
	UserId  int
	Balance float64
}

func CreateWallet(w Wallet) (Wallet, error) {
	db.Create(&w)
	return w, nil
}

func GetWallet(userId int)(Wallet, error) {
	var wallet Wallet
	db.FirstOrCreate(&wallet, Wallet{
		UserId: userId,
	})
	return wallet, nil
}

func UpdateWallet(wallet Wallet)(Wallet, error){
	db.Save(&wallet)
	return wallet, nil
}

func DeleteWallet(wallet Wallet) error {
	db.Delete(&wallet)
	return nil
}