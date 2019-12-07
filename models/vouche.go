package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type Vouch struct {
	Id        int
	VoucheeId string
	VoucherId string
}

func CreateVouch(vouch Vouch) (Vouch, error) {
	db.Create(&vouch)
	return vouch, nil
}

func GetVouch(id int)(Vouch, error){
	var vouch Vouch
	db.Find(&vouch, id)
	return vouch, nil
}


func GetAllVouchesForVouchee(voucheeId string)([]Vouch, error) {
	var vouches []Vouch
	if err := db.Where("vouchee_id = ?", voucheeId).Find(&vouches).Error; err != nil {
		return nil, err
	}

	return vouches, nil
}

func GetAllVouchesForVoucher(voucherId string)([]Vouch, error) {
	var vouches []Vouch
	if err := db.Where("voucher_id = ?", voucherId).Find(&vouches).Error; err != nil {
		return nil, err
	}

	return vouches, nil
}

func DeleteVouch(vouch Vouch) error {
	db.Delete(&vouch)
	return nil
}