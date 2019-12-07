package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

type Vouch struct {
	Id        int
	VoucheeId string
	VoucherId string
	VouchType string
	Amount float64
}

const (
	AutoPay = "AUTO_PAY"
	Default = "DEFAULT"
)

func CreateVouch(vouch Vouch) (Vouch, error) {
	db.Create(&vouch)
	return vouch, nil
}

func CreateVouches(vouches []Vouch) ([]Vouch, error) {
	var vchs []Vouch
	for _, v:= range vouches{
		db.Create(&v)
		vchs = append(vchs, v)
	}
	return vchs, nil
}

func GetVouch(id int)(Vouch, error){
	var vouch Vouch
	db.Find(&vouch, id)
	return vouch, nil
}

func IsVouch(voucheeId string, voucherId string, vouchType string) bool {
	var vouches []Vouch
	if err := db.Where("vouchee_id = ? AND vouch_type = ?", voucheeId, vouchType).Find(&vouches).Error; err != nil {
		return false
	}

	for _, v := range vouches {
		if v.VoucherId == voucherId {
			return true
		}
	}

	return false
}


func GetAllVouchesForVouchee(voucheeId string)([]Vouch, error) {
	var vouches []Vouch
	if err := db.Where("vouchee_id = ? AND vouch_type = ?", voucheeId, Default).Find(&vouches).Error; err != nil {
		return nil, err
	}

	return vouches, nil
}

func GetAllVouchesForVoucher(voucherId string)([]Vouch, error) {
	var vouches []Vouch
	if err := db.Where("voucher_id = ? AND vouch_type = ?", voucherId, Default).Find(&vouches).Error; err != nil {
		return nil, err
	}

	return vouches, nil
}

func DeleteVouch(vouch Vouch) error {
	db.Delete(&vouch)
	return nil
}