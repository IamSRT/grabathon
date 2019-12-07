package request_response

import "grabathon/models"

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

type Vouch struct {
	Id        int     `json:"id"`
	VoucheeId string  `json:"vouchee_id"`
	VoucherId string  `json:"voucher_id"`
	VouchType string  `json:"vouch_type"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
}

func GetVouchRequestResponse(vouch models.Vouch) Vouch {
	if vouch.VouchType == ""{
		vouch.VouchType = models.Default
	}

	vch := Vouch{
		Id:        vouch.Id,
		VoucheeId: vouch.VoucheeId,
		VoucherId: vouch.VoucherId,
		VouchType: vouch.VouchType,
		Status: vouch.Status,
		Amount: vouch.Amount,
	}

	return vch
}

func GetVouchModel(vouch Vouch) models.Vouch {
	if vouch.VouchType == ""{
		vouch.VouchType = models.Default
	}

	vch := models.Vouch{
		Id:        vouch.Id,
		VoucheeId: vouch.VoucheeId,
		VoucherId: vouch.VoucherId,
		VouchType: vouch.VouchType,
		Status: vouch.Status,
		Amount: vouch.Amount,
	}

	return vch
}