package request_response

import "grabathon/models"

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type Vouch struct {
	Id        int
	VoucheeId string
	VoucherId string
}

func GetVouchRequestResponse(vouch models.Vouch) Vouch {
	vch := Vouch{
		Id:        vouch.Id,
		VoucheeId: vouch.VoucheeId,
		VoucherId: vouch.VoucherId,
	}

	return vch
}

func GetVouchModel(vouch Vouch) models.Vouch {
	vch := models.Vouch{
		Id:        vouch.Id,
		VoucheeId: vouch.VoucheeId,
		VoucherId: vouch.VoucherId,
	}

	return vch
}