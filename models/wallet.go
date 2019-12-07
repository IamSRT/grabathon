package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type Wallet struct {
	Id      int
	UserId  string
	Balance float64
}

func CreateWallet(w Wallet) (Wallet, error) {
	if err := db.Create(&w).Error; err != nil {
		return Wallet{}, err
	}
	return w, nil
}

func GetWallet(userId string)(Wallet, error) {
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